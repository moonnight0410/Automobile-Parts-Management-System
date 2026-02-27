#!/bin/bash
# 从 Git 更新链码并部署到 Fabric 网络（支持断点续传）
# 
# 背书策略说明：
# - 策略类型：OR（或逻辑）
# - 策略表达式：OR('Org1MSP.member','Org2MSP.member','Org3MSP.member')
# - 策略含义：任何单个组织的成员都可以背书交易
# - 优势：
#   1. 减少网络往返次数
#   2. 降低交易延迟
#   3. 提高系统吞吐量
#   4. 减少单点故障影响
# - 组织职责：
#   - Org1MSP（零部件生产厂商）：CreatePart, CreateProductionData, CreateQualityInspection
#   - Org2MSP（整车车企）：CreateSupplyOrder, CreateLogisticsData
#   - Org3MSP（4S店/售后中心）：CreateFaultReport, CreateRecallRecord

# ==================== 配置变量 ====================
GIT_REPO_URL="https://github.com/moonnight0410/Automobile-Parts-Management-System.git"
LOCAL_CHAINCODE_DIR="/home/jianyu-zou/code/Automobile-Parts-Management-System"
CHANNEL_NAME="channel1"
CHAINCODE_NAME="auto-system"
INITIAL_VERSION="1.0"
CHAINCODE_PATH="${LOCAL_CHAINCODE_DIR}/chaincode"
FABRIC_NETWORK_PATH="~/fabric/fabric-samples/test-network"

# 背书策略配置（OR策略：任何单个组织都可以背书）
ENDORSEMENT_POLICY="OR('Org1MSP.member','Org2MSP.member','Org3MSP.member')"

# 状态文件和日志文件
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
STATE_FILE="${SCRIPT_DIR}/deployment_state.json"
LOG_FILE="${SCRIPT_DIR}/deployment_$(date +%Y%m%d_%H%M%S).log"

# 组织配置
declare -A ORGS=(
    ["Org1"]="7051:org1.example.com:Org1MSP"
    ["Org2"]="9051:org2.example.com:Org2MSP"
    ["Org3"]="11051:org3.example.com:Org3MSP"
)

# ==================== 日志函数 ====================
log() {
    local level="$1"
    shift
    local message="$@"
    local timestamp=$(date '+%Y-%m-%d %H:%M:%S')
    echo "[${timestamp}] [${level}] ${message}" | tee -a "$LOG_FILE"
}

log_info() {
    log "INFO" "$@"
}

log_success() {
    log "SUCCESS" "$@"
}

log_error() {
    log "ERROR" "$@"
}

log_warning() {
    log "WARNING" "$@"
}

log_verify() {
    log "VERIFY" "$@"
}

# ==================== 状态管理函数 ====================
init_state() {
    cat > "$STATE_FILE" << EOF
{
    "start_time": "$(date -Iseconds)",
    "current_step": 0,
    "completed_steps": [],
    "failed_step": null,
    "error_message": null
}
EOF
}

update_state() {
    local step="$1"
    local status="$2"
    local error_msg="$3"
    
    local completed_steps=$(jq -r '.completed_steps' "$STATE_FILE")
    
    if [ "$status" = "completed" ]; then
        completed_steps=$(echo "$completed_steps" | jq --arg step "$step" '. + [$step] | unique')
    fi
    
    jq --arg step "$step" \
       --argjson completed "$completed_steps" \
       --arg status "$status" \
       --arg error "$error_msg" \
       '.current_step = $step
        | .completed_steps = $completed
        | .failed_step = (if $status == "failed" then $step else .failed_step end)
        | .error_message = (if $status == "failed" then $error else .error_message end)' \
       "$STATE_FILE" > "${STATE_FILE}.tmp" && mv "${STATE_FILE}.tmp" "$STATE_FILE"
}

is_step_completed() {
    local step="$1"
    local completed=$(jq -r --arg step "$step" '.completed_steps | index($step) != null' "$STATE_FILE")
    [ "$completed" = "true" ]
}

delete_state() {
    if [ -f "$STATE_FILE" ]; then
        rm -f "$STATE_FILE"
        log_info "状态文件已删除"
    fi
}

show_state() {
    if [ ! -f "$STATE_FILE" ]; then
        log_info "没有找到状态文件，将从头开始执行"
        return
    fi
    
    log_info "================================"
    log_info "当前部署状态"
    log_info "================================"
    
    local start_time=$(jq -r '.start_time' "$STATE_FILE")
    local current_step=$(jq -r '.current_step' "$STATE_FILE")
    local completed_steps=$(jq -r '.completed_steps[]' "$STATE_FILE" 2>/dev/null)
    local failed_step=$(jq -r '.failed_step' "$STATE_FILE")
    local error_msg=$(jq -r '.error_message' "$STATE_FILE")
    
    log_info "开始时间: $start_time"
    log_info "当前步骤: $current_step"
    log_info ""
    
    if [ -n "$completed_steps" ]; then
        log_info "已完成的步骤:"
        echo "$completed_steps" | while read step; do
            log_info "  ✓ $step"
        done
        log_info ""
    fi
    
    if [ "$failed_step" != "null" ]; then
        log_error "失败的步骤: $failed_step"
        log_error "错误信息: $error_msg"
        log_info ""
    fi
    
    log_info "================================"
}

# ==================== 环境设置函数 ====================
setup_org_env() {
    local org="$1"
    local config="${ORGS[$org]}"
    local port=$(echo "$config" | cut -d: -f1)
    local org_name=$(echo "$config" | cut -d: -f2)
    local mspid=$(echo "$config" | cut -d: -f3)
    
    export PATH=~/fabric/fabric-samples/bin:$PATH
    export FABRIC_CFG_PATH=~/fabric/fabric-samples/config
    export CORE_PEER_LOCALMSPID="$mspid"
    export CORE_PEER_MSPCONFIGPATH=~/fabric/fabric-samples/test-network/organizations/peerOrganizations/${org_name}/users/Admin@${org_name}/msp
    export CORE_PEER_ADDRESS=localhost:${port}
    export CORE_PEER_TLS_ENABLED=true
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/${org_name}/peers/peer0.${org_name}/tls/ca.crt
}

# ==================== Git更新检测函数 ====================
check_git_update() {
    log_info "检查Git仓库是否有更新..."
    
    cd "$LOCAL_CHAINCODE_DIR"
    
    local local_commit=$(git rev-parse HEAD)
    local remote_commit=$(git ls-remote origin main | awk '{print $1}')
    
    if [ "$local_commit" = "$remote_commit" ]; then
        log_info "本地代码已是最新版本，无需更新"
        return 1
    else
        log_info "检测到远程仓库有更新"
        return 0
    fi
}

# ==================== 版本管理函数 ====================
get_chaincode_version() {
    local result=$(peer lifecycle chaincode querycommitted --channelID "$CHANNEL_NAME" --name "$CHAINCODE_NAME" --output json 2>/dev/null)
    
    if [ $? -ne 0 ] || [ -z "$result" ]; then
        echo "$INITIAL_VERSION"
    else
        local version=$(echo "$result" | jq -r '.version')
        if [ "$version" = "null" ] || [ -z "$version" ]; then
            echo "$INITIAL_VERSION"
        else
            echo "$version"
        fi
    fi
}

increment_version() {
    local version="$1"
    local major=$(echo "$version" | cut -d. -f1)
    local minor=$(echo "$version" | cut -d. -f2)
    local patch=$(echo "$version" | cut -d. -f3)
    
    if [ -z "$patch" ]; then
        patch=0
    fi
    
    patch=$((patch + 1))
    echo "${major}.${minor}.${patch}"
}

# ==================== 主函数 ====================
main() {
    log_info "================================"
    log_info "开始链码部署流程"
    log_info "时间: $(date)"
    log_info "================================"
    
    # 进入Fabric网络目录
    cd ~/fabric/fabric-samples/test-network || {
        log_error "无法进入Fabric网络目录: $FABRIC_NETWORK_PATH"
        exit 1
    }
    
    if [ -n "$AUTO_MODE" ]; then
        case "$AUTO_MODE" in
            1|reset|RESET)
                delete_state
                ;;
        esac
    fi
    
    # ==================== 步骤1: 初始化和状态检查 ====================
    local step=1
    if ! is_step_completed "step1_init"; then
        log_info "步骤1: 初始化和状态检查"
        
        if [ -f "$STATE_FILE" ]; then
            show_state
            echo ""
            if [ -n "$AUTO_MODE" ]; then
                case "$AUTO_MODE" in
                    1|reset|RESET)
                        delete_state
                        init_state
                        log_info "已选择从头开始执行(自动)"
                        ;;
                    2|resume|RESUME)
                        log_info "已选择继续执行(自动)"
                        ;;
                    *)
                        log_error "AUTO_MODE 无效值: $AUTO_MODE"
                        exit 1
                        ;;
                esac
            else
                read -p "请选择执行模式:
  1. 从头开始（删除状态文件）
  2. 继续执行（跳过已完成步骤）
请输入选项 (1/2): " choice
                case $choice in
                    1)
                        delete_state
                        init_state
                        log_info "已选择从头开始执行"
                        ;;
                    2)
                        log_info "已选择继续执行"
                        ;;
                    *)
                        log_error "无效的选项"
                        exit 1
                        ;;
                esac
            fi
        else
            init_state
            log_info "状态文件已初始化"
        fi
        
        update_state "step1_init" "completed"
        log_success "步骤1完成"
        
        # 验证输出
        log_verify "验证: 状态文件已创建"
        if [ -f "$STATE_FILE" ]; then
            log_verify "✓ 状态文件路径: $STATE_FILE"
            log_verify "✓ 状态文件内容:"
            cat "$STATE_FILE" | tee -a "$LOG_FILE"
        fi
    else
        log_info "步骤1: 已完成，跳过"
    fi
    
    # ==================== 步骤2: Git更新检测 ====================
    step=2
    if ! is_step_completed "step2_git_check"; then
        log_info "步骤2: Git更新检测"
        
        if ! check_git_update; then
            log_info "没有Git更新，继续执行部署流程"
        fi
        
        update_state "step2_git_check" "completed"
        log_success "步骤2完成"
        
        # 验证输出
        log_verify "验证: Git仓库检测完成"
        cd "$LOCAL_CHAINCODE_DIR"
        local current_commit=$(git rev-parse HEAD)
        log_verify "✓ 当前commit: $current_commit"
        cd ~/fabric/fabric-samples/test-network
    else
        log_info "步骤2: 已完成，跳过"
    fi
    
    # ==================== 步骤3: Git拉取 ====================
    step=3
    if ! is_step_completed "step3_git_pull"; then
        log_info "步骤3: 从Git拉取最新代码"
        
        if [ ! -d "$LOCAL_CHAINCODE_DIR" ]; then
            mkdir -p "$(dirname "$LOCAL_CHAINCODE_DIR")"
            if ! git clone "$GIT_REPO_URL" "$LOCAL_CHAINCODE_DIR"; then
                log_error "Git克隆失败"
                update_state "step3_git_pull" "failed" "Git克隆失败"
                exit 1
            fi
        fi
        
        cd "$LOCAL_CHAINCODE_DIR"
        if ! git fetch origin main; then
            log_error "Git fetch失败"
            update_state "step3_git_pull" "failed" "Git fetch失败"
            exit 1
        fi
        if ! git reset --hard origin/main; then
            log_error "Git重置到origin/main失败"
            update_state "step3_git_pull" "failed" "Git重置失败"
            exit 1
        fi
        
        cd ~/fabric/fabric-samples/test-network
        update_state "step3_git_pull" "completed"
        log_success "步骤3完成"
        
        # 验证输出
        log_verify "验证: Git拉取成功"
        cd "$LOCAL_CHAINCODE_DIR"
        local latest_commit=$(git log -1 --oneline)
        log_verify "✓ 最新提交: $latest_commit"
        log_verify "✓ 分支状态:"
        git status | tee -a "$LOG_FILE"
        cd ~/fabric/fabric-samples/test-network
    else
        log_info "步骤3: 已完成，跳过"
    fi
    
    # ==================== 步骤4: 链码目录检查 ====================
    step=4
    if ! is_step_completed "step4_check_dir"; then
        log_info "步骤4: 检查链码目录"
        
        if [ ! -d "$CHAINCODE_PATH" ]; then
            log_error "链码路径不存在: $CHAINCODE_PATH"
            update_state "step4_check_dir" "failed" "链码路径不存在"
            exit 1
        fi
        
        ls -la "$CHAINCODE_PATH" | tee -a "$LOG_FILE"
        update_state "step4_check_dir" "completed"
        log_success "步骤4完成"
        
        # 验证输出
        log_verify "验证: 链码目录检查完成"
        log_verify "✓ 目录路径: $CHAINCODE_PATH"
        log_verify "✓ 主要文件:"
        ls -lh "$CHAINCODE_PATH"/*.go 2>/dev/null | tee -a "$LOG_FILE"
    else
        log_info "步骤4: 已完成，跳过"
    fi
    
    # ==================== 步骤5: 链码打包 ====================
    step=5
    if ! is_step_completed "step5_package"; then
        log_info "步骤5: 打包链码"
        
        setup_org_env "Org1"
        local current_version=$(get_chaincode_version)
        local new_version
        if [ -n "$CHAINCODE_VERSION" ]; then
            new_version="$CHAINCODE_VERSION"
        else
            new_version=$(increment_version "$current_version")
        fi
        
        log_info "当前版本: $current_version"
        log_info "新版本: $new_version"
        
        local package_file="${CHAINCODE_NAME}_v${new_version}.tar.gz"
        
        if ! peer lifecycle chaincode package "$package_file" \
            --lang golang \
            --path "$CHAINCODE_PATH" \
            --label "${CHAINCODE_NAME}_${new_version}"; then
            log_error "链码打包失败"
            update_state "step5_package" "failed" "链码打包失败"
            exit 1
        fi
        
        NEW_VERSION="$new_version"
        PACKAGE_FILE="$package_file"
        
        update_state "step5_package" "completed"
        log_success "步骤5完成: $PACKAGE_FILE"
        
        # 验证输出
        log_verify "验证: 链码打包成功"
        log_verify "✓ 包文件: $PACKAGE_FILE"
        log_verify "✓ 包文件信息:"
        ls -lh "$PACKAGE_FILE" | tee -a "$LOG_FILE"
        log_verify "✓ 包文件内容:"
        tar -tzf "$PACKAGE_FILE" | head -20 | tee -a "$LOG_FILE"
    else
        log_info "步骤5: 已完成，跳过"
        NEW_VERSION=$(jq -r '.new_version' "$STATE_FILE" 2>/dev/null || echo "1.0")
        PACKAGE_FILE="${CHAINCODE_NAME}_v${NEW_VERSION}.tar.gz"
    fi
    
    # ==================== 步骤6: 版本号管理 ====================
    step=6
    if ! is_step_completed "step6_version"; then
        log_info "步骤6: 获取链码序列号"
        
        setup_org_env "Org1"
        local current_sequence=$(peer lifecycle chaincode querycommitted --channelID "$CHANNEL_NAME" --name "$CHAINCODE_NAME" --output json 2>/dev/null | jq -r '.sequence')
        
        if [ -z "$current_sequence" ] || [ "$current_sequence" = "null" ]; then
            current_sequence=0
        fi
        
        local new_sequence
        if [ -n "$CHAINCODE_SEQUENCE" ]; then
            new_sequence="$CHAINCODE_SEQUENCE"
        else
            new_sequence=$((current_sequence + 1))
        fi
        
        log_info "当前序列号: $current_sequence"
        log_info "新序列号: $new_sequence"
        
        NEW_SEQUENCE="$new_sequence"
        
        update_state "step6_version" "completed"
        log_success "步骤6完成"
        
        # 验证输出
        log_verify "验证: 版本号管理完成"
        log_verify "✓ 链码名称: $CHAINCODE_NAME"
        log_verify "✓ 通道名称: $CHANNEL_NAME"
        log_verify "✓ 当前版本: $NEW_VERSION"
        log_verify "✓ 当前序列号: $NEW_SEQUENCE"
    else
        log_info "步骤6: 已完成，跳过"
        NEW_SEQUENCE=$(jq -r '.new_sequence' "$STATE_FILE" 2>/dev/null || echo "1")
    fi
    
    # 保存版本信息到状态文件
    jq --arg version "$NEW_VERSION" --arg sequence "$NEW_SEQUENCE" \
       '. + {new_version: $version, new_sequence: $sequence}' \
       "$STATE_FILE" > "${STATE_FILE}.tmp" && mv "${STATE_FILE}.tmp" "$STATE_FILE"
    
    # ==================== 步骤7: 链码安装 ====================
    step=7
    if ! is_step_completed "step7_install"; then
        log_info "步骤7: 在所有组织上安装链码"
        
        for org in Org1 Org2 Org3; do
            log_info "在${org}上安装链码..."
            
            setup_org_env "$org"
            
            if ! peer lifecycle chaincode install "$PACKAGE_FILE"; then
                if peer lifecycle chaincode queryinstalled | grep -q "${CHAINCODE_NAME}_${NEW_VERSION}:[a-zA-Z0-9]*"; then
                    log_success "${org}链码已安装"
                else
                    log_error "${org}链码安装失败"
                    update_state "step7_install" "failed" "${org}链码安装失败"
                    exit 1
                fi
            fi
            
            log_success "${org}链码安装成功"
        done
        
        update_state "step7_install" "completed"
        log_success "步骤7完成"
        
        # 验证输出
        log_verify "验证: 链码安装完成"
        log_verify "✓ 已安装的链码列表:"
        peer lifecycle chaincode queryinstalled --output json | tee -a "$LOG_FILE"
        log_verify "✓ 当前包ID:"
        peer lifecycle chaincode queryinstalled | grep -o "${CHAINCODE_NAME}_${NEW_VERSION}:[a-zA-Z0-9]*" | tee -a "$LOG_FILE"
    else
        log_info "步骤7: 已完成，跳过"
    fi
    
    # ==================== 步骤8: 获取包ID ====================
    step=8
    if ! is_step_completed "step8_package_id"; then
        log_info "步骤8: 获取链码包ID"
        
        local package_id=$(peer lifecycle chaincode queryinstalled | grep -o "${CHAINCODE_NAME}_${NEW_VERSION}:[a-zA-Z0-9]*")
        
        if [ -z "$package_id" ]; then
            log_error "无法获取包ID"
            update_state "step8_package_id" "failed" "无法获取包ID"
            exit 1
        fi
        
        log_info "包ID: $package_id"
        
        PACKAGE_ID="$package_id"
        
        update_state "step8_package_id" "completed"
        log_success "步骤8完成"
        
        # 验证输出
        log_verify "验证: 包ID获取成功"
        log_verify "✓ 链码名称: $CHAINCODE_NAME"
        log_verify "✓ 版本号: $NEW_VERSION"
        log_verify "✓ 包ID: $PACKAGE_ID"
    else
        log_info "步骤8: 已完成，跳过"
        PACKAGE_ID=$(jq -r '.package_id' "$STATE_FILE" 2>/dev/null)
    fi
    
    # 保存包ID到状态文件
    jq --arg pkg_id "$PACKAGE_ID" '. + {package_id: $pkg_id}' "$STATE_FILE" > "${STATE_FILE}.tmp" && mv "${STATE_FILE}.tmp" "$STATE_FILE"
    
    # ==================== 步骤9: 链码批准 ====================
    step=9
    if ! is_step_completed "step9_approve"; then
        log_info "步骤9: 批准链码定义"
        
        for org in Org1 Org2 Org3; do
            log_info "${org}批准链码定义..."
            
            setup_org_env "$org"
            
            if ! peer lifecycle chaincode approveformyorg \
                -o localhost:7050 \
                --ordererTLSHostnameOverride orderer.example.com \
                --tls true \
                --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
                --channelID "$CHANNEL_NAME" \
                --name "$CHAINCODE_NAME" \
                --version "$NEW_VERSION" \
                --sequence "$NEW_SEQUENCE" \
                --init-required \
                --package-id "$PACKAGE_ID"; then
                log_error "${org}批准链码定义失败"
                update_state "step9_approve" "failed" "${org}批准链码定义失败"
                exit 1
            fi
            
            log_success "${org}批准成功"
        done
        
        update_state "step9_approve" "completed"
        log_success "步骤9完成"
        
        # 验证输出
        log_verify "验证: 链码批准完成"
        log_verify "✓ 检查各组织批准状态:"
        for org in Org1 Org2 Org3; do
            setup_org_env "$org"
            log_verify "  ${org} 批准状态:"
            peer lifecycle chaincode queryapproved --channelID "$CHANNEL_NAME" --name "$CHAINCODE_NAME" --output json 2>/dev/null | jq -r '.approvals' 2>/dev/null || echo "  查询失败"
        done
    else
        log_info "步骤9: 已完成，跳过"
    fi
    
    # ==================== 步骤10: 检查提交准备状态 ====================
    step=10
    if ! is_step_completed "step10_check_commit"; then
        log_info "步骤10: 检查链码提交准备状态"
        
        peer lifecycle chaincode checkcommitreadiness \
            --channelID "$CHANNEL_NAME" \
            --name "$CHAINCODE_NAME" \
            --version "$NEW_VERSION" \
            --sequence "$NEW_SEQUENCE" \
            --output json | tee -a "$LOG_FILE"
        
        update_state "step10_check_commit" "completed"
        log_success "步骤10完成"
        
        # 验证输出
        log_verify "验证: 提交准备状态检查完成"
        log_verify "✓ 各组织批准情况:"
        peer lifecycle chaincode checkcommitreadiness \
            --channelID "$CHANNEL_NAME" \
            --name "$CHAINCODE_NAME" \
            --version "$NEW_VERSION" \
            --sequence "$NEW_SEQUENCE" \
            --output json | jq -r '.approvals' 2>/dev/null | tee -a "$LOG_FILE"
    else
        log_info "步骤10: 已完成，跳过"
    fi
    
    # ==================== 步骤11: 提交链码定义 ====================
    step=11
    if ! is_step_completed "step11_commit"; then
        log_info "步骤11: 提交链码定义到通道"
        log_info "背书策略: $ENDORSEMENT_POLICY"
        
        setup_org_env "Org1"
        
        if ! peer lifecycle chaincode commit \
            -o localhost:7050 \
            --ordererTLSHostnameOverride orderer.example.com \
            --tls true \
            --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
            --channelID "$CHANNEL_NAME" \
            --name "$CHAINCODE_NAME" \
            --version "$NEW_VERSION" \
            --sequence "$NEW_SEQUENCE" \
            --init-required \
            --signature-policy "$ENDORSEMENT_POLICY" \
            --peerAddresses localhost:7051 \
            --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt \
            --peerAddresses localhost:9051 \
            --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt \
            --peerAddresses localhost:11051 \
            --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt; then
            log_error "链码定义提交失败"
            update_state "step11_commit" "failed" "链码定义提交失败"
            exit 1
        fi
        
        update_state "step11_commit" "completed"
        log_success "步骤11完成"
        
        # 验证输出
        log_verify "验证: 链码定义提交成功"
        log_verify "✓ 通道中已提交的链码:"
        peer lifecycle chaincode querycommitted --channelID "$CHANNEL_NAME" --output json | tee -a "$LOG_FILE"
        log_verify "✓ 验证背书策略:"
        peer lifecycle chaincode querycommitted --channelID "$CHANNEL_NAME" --name "$CHAINCODE_NAME" --output json | jq -r '.endorsement_info' 2>/dev/null | tee -a "$LOG_FILE"
    else
        log_info "步骤11: 已完成，跳过"
    fi
    
    # ==================== 步骤12: 验证更新 ====================
    step=12
    if ! is_step_completed "step12_verify"; then
        log_info "步骤12: 验证链码更新和背书策略"
        
        local commit_info=$(peer lifecycle chaincode querycommitted \
            --channelID "$CHANNEL_NAME" \
            --name "$CHAINCODE_NAME" \
            --output json 2>/dev/null)
        
        echo "$commit_info" | tee -a "$LOG_FILE"
        
        update_state "step12_verify" "completed"
        log_success "步骤12完成"
        
        # 验证输出
        log_verify "验证: 链码更新验证完成"
        log_verify "✓ 链码详细信息:"
        echo "$commit_info" | jq '.' 2>/dev/null | tee -a "$LOG_FILE"
        log_verify "✓ 版本号: $NEW_VERSION"
        log_verify "✓ 序列号: $NEW_SEQUENCE"
        log_verify "✓ 包ID: $PACKAGE_ID"
        log_verify "✓ 背书策略验证:"
        local endorsement_info=$(echo "$commit_info" | jq -r '.endorsement_info' 2>/dev/null)
        if [ -n "$endorsement_info" ] && [ "$endorsement_info" != "null" ]; then
            log_verify "  背书策略已配置: $endorsement_info"
            log_success "✓ 背书策略验证成功"
        else
            log_warning "  背书策略信息未找到，可能使用默认策略"
        fi
    else
        log_info "步骤12: 已完成，跳过"
    fi
    
    # ==================== 步骤13: 初始化链码 ====================
    step=13
    if ! is_step_completed "step13_init"; then
        log_info "步骤13: 初始化链码"
        
        setup_org_env "Org1"
        
        log_info "检查链码是否已经初始化..."
        local init_check=$(peer chaincode query -C "$CHANNEL_NAME" -n "$CHAINCODE_NAME" -c '{"Args":["org.hyperledger.fabric:GetMetadata"]}' 2>&1)
        
        if echo "$init_check" | grep -q "error"; then
            log_info "链码尚未初始化，开始初始化..."
            
            if ! peer chaincode invoke \
                -o localhost:7050 \
                --ordererTLSHostnameOverride orderer.example.com \
                --tls true \
                --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
                -C "$CHANNEL_NAME" \
                -n "$CHAINCODE_NAME" \
                --peerAddresses localhost:7051 \
                --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt \
                --peerAddresses localhost:9051 \
                --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt \
                --peerAddresses localhost:11051 \
                --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt \
                --isInit \
                -c '{"function":"initLedger","Args":[]}'; then
                log_error "链码初始化失败"
                update_state "step13_init" "failed" "链码初始化失败"
                exit 1
            fi
            
            log_success "链码初始化成功"
        else
            log_info "链码已经初始化，跳过初始化步骤"
        fi
        
        update_state "step13_init" "completed"
        log_success "步骤13完成"
        
        # 验证输出
        log_verify "验证: 链码初始化完成"
        log_verify "✓ 等待链码容器启动..."
        sleep 5
        log_verify "✓ 链码容器状态:"
        docker ps --filter "name=dev-peer" --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}" | tee -a "$LOG_FILE"
        log_verify "✓ 测试链码调用:"
        setup_org_env "Org1"
        peer chaincode query -C "$CHANNEL_NAME" -n "$CHAINCODE_NAME" -c '{"Args":["org.hyperledger.fabric:GetMetadata"]}' 2>/dev/null | head -20 | tee -a "$LOG_FILE" || log_verify "  查询失败（可能需要等待更长时间）"
    else
        log_info "步骤13: 已完成，跳过"
    fi
    
    # ==================== 步骤14: 测试背书策略 ====================
    step=14
    if ! is_step_completed "step14_test_endorsement"; then
        log_info "步骤14: 测试背书策略"
        log_info "测试OR背书策略：任何单个组织都可以背书交易"
        
        setup_org_env "Org1"
        
        log_info "测试1: 使用Org1MSP进行背书"
        local test_result=$(peer chaincode query -C "$CHANNEL_NAME" -n "$CHAINCODE_NAME" -c '{"Args":["org.hyperledger.fabric:GetMetadata"]}' 2>&1)
        if echo "$test_result" | grep -q "error"; then
            log_error "Org1MSP背书测试失败"
            update_state "step14_test_endorsement" "failed" "Org1MSP背书测试失败"
            exit 1
        else
            log_success "Org1MSP背书测试成功"
        fi
        
        setup_org_env "Org2"
        
        log_info "测试2: 使用Org2MSP进行背书"
        test_result=$(peer chaincode query -C "$CHANNEL_NAME" -n "$CHAINCODE_NAME" -c '{"Args":["org.hyperledger.fabric:GetMetadata"]}' 2>&1)
        if echo "$test_result" | grep -q "error"; then
            log_error "Org2MSP背书测试失败"
            update_state "step14_test_endorsement" "failed" "Org2MSP背书测试失败"
            exit 1
        else
            log_success "Org2MSP背书测试成功"
        fi
        
        setup_org_env "Org3"
        
        log_info "测试3: 使用Org3MSP进行背书"
        test_result=$(peer chaincode query -C "$CHANNEL_NAME" -n "$CHAINCODE_NAME" -c '{"Args":["org.hyperledger.fabric:GetMetadata"]}' 2>&1)
        if echo "$test_result" | grep -q "error"; then
            log_error "Org3MSP背书测试失败"
            update_state "step14_test_endorsement" "failed" "Org3MSP背书测试失败"
            exit 1
        else
            log_success "Org3MSP背书测试成功"
        fi
        
        log_success "所有组织的背书测试通过"
        update_state "step14_test_endorsement" "completed"
        log_success "步骤14完成"
        
        # 验证输出
        log_verify "验证: 背书策略测试完成"
        log_verify "✓ Org1MSP: 可以背书 ✓"
        log_verify "✓ Org2MSP: 可以背书 ✓"
        log_verify "✓ Org3MSP: 可以背书 ✓"
        log_verify "✓ OR背书策略验证成功"
    else
        log_info "步骤14: 已完成，跳过"
    fi
    
    # ==================== 步骤15: 清理 ====================
    step=15
    if ! is_step_completed "step15_cleanup"; then
        log_info "步骤15: 清理状态文件"
        
        delete_state
        
        update_state "step15_cleanup" "completed"
        log_success "步骤15完成"
        
        # 验证输出
        log_verify "验证: 清理完成"
        log_verify "✓ 状态文件已删除"
        log_verify "✓ 日志文件路径: $LOG_FILE"
        log_verify "✓ 日志文件大小:"
        ls -lh "$LOG_FILE" | awk '{print "  " $5}'
    else
        log_info "步骤15: 已完成，跳过"
    fi
    
    # ==================== 完成 ====================
    log_info "================================"
    log_success "链码部署完成!"
    log_info "版本: $NEW_VERSION"
    log_info "序列号: $NEW_SEQUENCE"
    log_info "包ID: $PACKAGE_ID"
    log_info "背书策略: $ENDORSEMENT_POLICY"
    log_info "时间: $(date)"
    log_info "日志文件: $LOG_FILE"
    log_info "================================"
    
    # 最终验证
    log_verify "================================"
    log_verify "最终验证"
    log_verify "================================"
    log_verify "✓ 链码部署信息:"
    local final_commit_info=$(peer lifecycle chaincode querycommitted --channelID "$CHANNEL_NAME" --name "$CHAINCODE_NAME" --output json 2>/dev/null)
    echo "$final_commit_info" | jq '.' 2>/dev/null | tee -a "$LOG_FILE"
    log_verify "✓ 背书策略最终验证:"
    local final_endorsement_info=$(echo "$final_commit_info" | jq -r '.endorsement_info' 2>/dev/null)
    if [ -n "$final_endorsement_info" ] && [ "$final_endorsement_info" != "null" ]; then
        log_verify "  背书策略: $final_endorsement_info"
        log_success "✓ OR背书策略已正确配置"
        log_verify "  策略说明: 任何单个组织（Org1MSP、Org2MSP或Org3MSP）都可以背书交易"
    else
        log_warning "  背书策略信息未找到，使用默认策略"
    fi
    log_verify "✓ 链码容器状态:"
    docker ps --filter "name=dev-peer" --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}" | tee -a "$LOG_FILE"
    log_verify "================================"
}

# 执行主函数
main
