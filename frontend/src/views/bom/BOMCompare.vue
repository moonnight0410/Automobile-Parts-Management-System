<!--
  BOM比较页面
-->

<template>
  <div class="bom-compare">
    <div class="page-header">
      <h2>BOM比较</h2>
    </div>
    
    <a-card :bordered="false">
      <a-form layout="inline">
        <a-form-item label="生产BOM">
          <a-select placeholder="选择生产BOM" style="width: 200px" />
        </a-form-item>
        <a-form-item label="研发BOM">
          <a-select placeholder="选择研发BOM" style="width: 200px" />
        </a-form-item>
        <a-form-item>
          <a-button type="primary">开始比较</a-button>
        </a-form-item>
      </a-form>
      
      <a-divider />
      
      <a-empty description="请选择两个BOM进行比较" />
    </a-card>
  </div>
</template>

<style scoped>
.bom-compare { padding: 0; }
.page-header { margin-bottom: 24px; }
.page-header h2 { font-size: 20px; font-weight: 600; margin: 0; }
</style>
