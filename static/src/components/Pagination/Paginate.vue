<!--
如何使用？
1.导入并添加components, import ToolBar from '~/components/ToolBar/ToolBar.vue',
2.绑定你的api、参数和数据处理函数，会自动生成分页
<Paginate
    api="postList"
    :params="params"
    :refresh="refresh"
    @val-change="PaginateData"
    >
</Paginate>
3.当你再次改动参数时，执行 refresh=!refresh 即可刷新。用于搜索。

-->

<template>
  <div class="pagination">
    <div class="total">共 {{total}} 条数据，当页有 {{currentDataLength}} 条数据。</div>
    <el-pagination
      background
      @current-change="change"
      :page-sizes="[10, 15, 20, 50]"
      :page-size="$Config.pageSize"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total">
    </el-pagination>
  </div>
</template>
<script>
  export default {
    name: 'Paginate',
    props: {
      api: Function,
      params: Object,
      refresh: Boolean,
    },
    data() {
      return {
        total: 0,
        current: 0,
        currentDataLength: 0,
      }
    },
    methods: {
      paginate: function () {
        let params = this.params;
        params.per_page = this.$Config.pageSize;
        params.page = this.current;
        let thisApp = this;

        thisApp.api(params, function (resp) {
          thisApp.total = parseInt(resp.headers['x-pagination-total']);
          thisApp.currentDataLength = (resp.data).length
          thisApp.$emit('val-change', resp.data);
        })
      },
      change: function (page) {
        this.current = page;
        this.paginate();
      }
    },
    watch: {

      'refresh': {
        handler: function () {
          this.current = 0;
          this.paginate();
        },
        // 深度观察
        deep: true
      }
    },
    mounted: function () {
      this.paginate();
    },
  }
</script>
<style lang="less" scoped>
  .pagination {
    padding: 7px;
    border: 1px solid #ebeef5;
    background: #fff;
    text-align: right;
  }
  .total {
    float: left;
    margin-top: 8px;
    margin-left: 10px;
    font-size: 13px;
    color: #606266;
  }
</style>
