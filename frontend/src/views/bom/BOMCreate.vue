<!--
  BOM创建页面
-->

<template>
  <div class="bom-create">
    <div class="page-header">
      <h2>创建BOM</h2>
      <a-button @click="$router.back()">返回</a-button>
    </div>
    
    <a-card :bordered="false">
      <a-form :label-col="{ span: 4 }" :wrapper-col="{ span: 16 }">
        <a-form-item label="BOM ID" required>
          <a-input placeholder="请输入BOM ID" />
        </a-form-item>
        <a-form-item label="BOM类型" required>
          <a-select placeholder="请选择类型">
            <a-select-option value="生产BOM">生产BOM</a-select-option>
            <a-select-option value="研发BOM">研发BOM</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="车型" required>
          <a-input placeholder="请输入车型" />
        </a-form-item>
        <a-form-item label="版本号" required>
          <a-input placeholder="请输入版本号" />
        </a-form-item>
        <a-form-item :wrapper-col="{ offset: 4 }">
          <a-button type="primary">提交</a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>

<style scoped>
.bom-create { padding: 0; }
.page-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 24px; }
.page-header h2 { font-size: 20px; font-weight: 600; margin: 0; }
</style>
