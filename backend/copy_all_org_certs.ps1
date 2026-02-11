# Copy all Fabric organization certificates from VM to local
# Used to copy certificates for org1, org2, and org3

# ==================== Configuration ====================

$vm_ip = "192.168.220.129"
$vm_user = "jianyu-zou"
$base_fabric_path = "~/fabric/fabric-samples/test-network/organizations/peerOrganizations"
$local_base_dir = "C:\Users\zjy20\Desktop\code\Automobile-Parts-Management-System\backend\fabric-certs"

# Organizations to copy
$organizations = @("org1", "org2", "org3")

# ==================== Functions ====================

function Copy-OrgCertificates {
    param([string]$OrgName)
    
    $org_path = "$base_fabric_path/$OrgName.example.com"
    $local_dir = "$local_base_dir\$OrgName"
    
    Write-Host ""
    Write-Host "========================================" -ForegroundColor Cyan
    Write-Host "  Processing $OrgName" -ForegroundColor Cyan
    Write-Host "========================================" -ForegroundColor Cyan
    
    # Check if org exists on VM
    $result = ssh "${vm_user}@${vm_ip}" "test -d $org_path && echo 'EXISTS' || echo 'NOT FOUND'" 2>&1
    if ($result -ne "EXISTS") {
        Write-Host "[Skip] $OrgName not found on VM, skipping..." -ForegroundColor Yellow
        return $false
    }
    
    Write-Host "[Found] $OrgName exists on VM" -ForegroundColor Green
    
    # Create local directory
    if (Test-Path $local_dir) {
        Write-Host "[Hint] Directory already exists: $local_dir" -ForegroundColor Yellow
        Remove-Item -Path $local_dir -Recurse -Force
        Write-Host "[Deleted] Old directory removed" -ForegroundColor Yellow
    }
    
    New-Item -ItemType Directory -Path "$local_dir\signcerts" -Force | Out-Null
    New-Item -ItemType Directory -Path "$local_dir\keystore" -Force | Out-Null
    New-Item -ItemType Directory -Path "$local_dir\tls" -Force | Out-Null
    Write-Host "[Success] Certificate directory created: $local_dir" -ForegroundColor Green
    
    # Copy Admin certificate
    Write-Host ""
    Write-Host "[Copy] Copying Admin certificate..." -ForegroundColor Cyan
    $adminCertRemote = "$org_path/users/Admin@$OrgName.example.com/msp/signcerts/Admin@$OrgName.example.com-cert.pem"
    $adminCertLocal = "$local_dir\signcerts\cert.pem"
    ssh "${vm_user}@${vm_ip}" "cat $adminCertRemote" > $adminCertLocal 2>&1
    
    if (Test-Path $adminCertLocal) {
        $fileSize = (Get-Item $adminCertLocal).Length
        Write-Host "[Success] Admin certificate copied (size: $fileSize bytes)" -ForegroundColor Green
    } else {
        Write-Host "[Failed] Admin certificate copy failed" -ForegroundColor Red
        return $false
    }
    
    # Copy Admin private key
    Write-Host "[Copy] Copying Admin private key..." -ForegroundColor Cyan
    $adminKeyRemote = "$org_path/users/Admin@$OrgName.example.com/msp/keystore/priv_sk"
    $adminKeyLocal = "$local_dir\keystore\key.pem"
    ssh "${vm_user}@${vm_ip}" "cat $adminKeyRemote" > $adminKeyLocal 2>&1
    
    if (Test-Path $adminKeyLocal) {
        $fileSize = (Get-Item $adminKeyLocal).Length
        Write-Host "[Success] Admin private key copied (size: $fileSize bytes)" -ForegroundColor Green
    } else {
        Write-Host "[Failed] Admin private key copy failed" -ForegroundColor Red
        return $false
    }
    
    # Copy TLS certificate
    Write-Host "[Copy] Copying TLS certificate..." -ForegroundColor Cyan
    $tlsCertRemote = "$org_path/peers/peer0.$OrgName.example.com/tls/ca.crt"
    $tlsCertLocal = "$local_dir\tls\ca.crt"
    ssh "${vm_user}@${vm_ip}" "cat $tlsCertRemote" > $tlsCertLocal 2>&1
    
    if (Test-Path $tlsCertLocal) {
        $fileSize = (Get-Item $tlsCertLocal).Length
        Write-Host "[Success] TLS certificate copied (size: $fileSize bytes)" -ForegroundColor Green
    } else {
        Write-Host "[Failed] TLS certificate copy failed" -ForegroundColor Red
        return $false
    }
    
    # Verify certificate files
    Write-Host ""
    Write-Host "[Verify] Verifying certificate files..." -ForegroundColor Cyan
    
    $allValid = $true
    
    if (Test-Path $adminCertLocal) {
        $certContent = Get-Content $adminCertLocal -Raw
        if ($certContent -match "-----BEGIN CERTIFICATE-----" -and $certContent -match "-----END CERTIFICATE-----") {
            Write-Host "[Valid] Admin certificate file" -ForegroundColor Green
        } else {
            Write-Host "[Invalid] Admin certificate file format error" -ForegroundColor Red
            $allValid = $false
        }
    } else {
        Write-Host "[Missing] Admin certificate file not found" -ForegroundColor Red
        $allValid = $false
    }
    
    if (Test-Path $adminKeyLocal) {
        $keyContent = Get-Content $adminKeyLocal -Raw
        if ($keyContent -match "-----BEGIN PRIVATE KEY-----" -or $keyContent -match "-----BEGIN EC PRIVATE KEY-----") {
            Write-Host "[Valid] Admin private key file" -ForegroundColor Green
        } else {
            Write-Host "[Invalid] Admin private key file format error" -ForegroundColor Red
            $allValid = $false
        }
    } else {
        Write-Host "[Missing] Admin private key file not found" -ForegroundColor Red
        $allValid = $false
    }
    
    if (Test-Path $tlsCertLocal) {
        $tlsContent = Get-Content $tlsCertLocal -Raw
        if ($tlsContent -match "-----BEGIN CERTIFICATE-----" -and $tlsContent -match "-----END CERTIFICATE-----") {
            Write-Host "[Valid] TLS certificate file" -ForegroundColor Green
        } else {
            Write-Host "[Invalid] TLS certificate file format error" -ForegroundColor Red
            $allValid = $false
        }
    } else {
        Write-Host "[Missing] TLS certificate file not found" -ForegroundColor Red
        $allValid = $false
    }
    
    return $allValid
}

# ==================== Main Program ====================

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Fabric Certificate Copy Tool v2.0" -ForegroundColor Cyan
Write-Host "  Copy certificates for all organizations" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

Write-Host "Configuration:" -ForegroundColor Yellow
Write-Host "  VM IP: $vm_ip" -ForegroundColor White
Write-Host "  VM User: $vm_user" -ForegroundColor White
Write-Host "  Base Fabric Path: $base_fabric_path" -ForegroundColor White
Write-Host "  Local Base Dir: $local_base_dir" -ForegroundColor White
Write-Host "  Organizations: $($organizations -join ', ')" -ForegroundColor White
Write-Host ""

# Check SSH connection
Write-Host "[Test] Checking SSH connection to $vm_ip..." -ForegroundColor Cyan
$result = ssh -o ConnectTimeout=10 -o BatchMode=yes "${vm_user}@${vm_ip}" "echo 'Connection successful'" 2>&1
if ($LASTEXITCODE -eq 0) {
    Write-Host "[Success] SSH connection is OK (no password required)" -ForegroundColor Green
} else {
    Write-Host "[Failed] SSH connection failed" -ForegroundColor Red
    exit 1
}

# Copy certificates for each organization
$results = @()

foreach ($org in $organizations) {
    $success = Copy-OrgCertificates -OrgName $org
    $results += [PSCustomObject]@{
        OrgName = $org
        Success = $success
    }
}

# Show summary
Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "          Copy Summary" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

$successCount = 0
$failCount = 0

foreach ($result in $results) {
    if ($result.Success) {
        Write-Host "[Success] $($result.OrgName): All certificate files copied and verified!" -ForegroundColor Green
        $successCount++
    } else {
        Write-Host "[Failed] $($result.OrgName): Certificate copy process encountered errors" -ForegroundColor Red
        $failCount++
    }
}

Write-Host ""
Write-Host "Total organizations: $($results.Count)" -ForegroundColor Cyan
Write-Host "Successful: $successCount" -ForegroundColor Green
Write-Host "Failed: $failCount" -ForegroundColor Red

if ($failCount -eq 0) {
    Write-Host ""
    Write-Host "All certificates copied successfully!" -ForegroundColor Green
    Write-Host ""
    Write-Host "Next steps:" -ForegroundColor Yellow
    Write-Host "1. Update environment variable configuration" -ForegroundColor White
    Write-Host "2. Run test script to verify Fabric connection" -ForegroundColor White
    Write-Host "3. Integrate into existing service" -ForegroundColor White
    exit 0
} else {
    Write-Host ""
    Write-Host "Some certificate copies failed. Please check the errors above." -ForegroundColor Yellow
    exit 1
}
