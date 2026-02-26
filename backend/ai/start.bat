@echo off
echo ========================================
echo AI服务启动脚本
echo ========================================
echo.

REM 检查Python是否安装
python --version >nul 2>&1
if %errorlevel% neq 0 (
    echo 错误: 未找到Python，请先安装Python 3.8或更高版本
    pause
    exit /b 1
)

echo 检测到Python版本:
python --version
echo.

REM 检查虚拟环境是否存在
if not exist "venv" (
    echo 创建虚拟环境...
    python -m venv venv
    if %errorlevel% neq 0 (
        echo 错误: 创建虚拟环境失败
        pause
        exit /b 1
    )
    echo 虚拟环境创建成功
    echo.
)

REM 激活虚拟环境
echo 激活虚拟环境...
call venv\Scripts\activate.bat
if %errorlevel% neq 0 (
    echo 错误: 激活虚拟环境失败
    pause
    exit /b 1
)
echo 虚拟环境激活成功
echo.

REM 安装依赖
echo 检查并安装依赖...
pip install -r requirements.txt
if %errorlevel% neq 0 (
    echo 警告: 安装依赖时出现错误，但继续启动服务
)
echo.

REM 启动AI服务
echo ========================================
echo 启动AI服务...
echo ========================================
echo.
echo 服务地址: http://localhost:5000
echo 健康检查: http://localhost:5000/health
echo.
echo 按 Ctrl+C 停止服务
echo.

python app.py

pause
