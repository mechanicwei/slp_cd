<template>
  <div class="role-list">
    <ToolBar>
      <el-button type="primary" icon="el-icon-arrow-left" size="small" @click="$router.push('/deploy_repos')">返回</el-button>
    </ToolBar>
    <el-table
        :data="recordsData"
        border
        ref="table"
        style="width: 100%">
      <el-table-column
        label="部署者"
        width="180">
        <template slot-scope="scope">
          <el-popover trigger="hover" placement="top">
            <a :href="scope.row.deploy_user.github_url" target="_blank">
              {{ scope.row.deploy_user.name }}
            </a>
            <div slot="reference" class="name-wrapper">
              <img :src="scope.row.deploy_user.avatar_url" :alt="scope.row.deploy_user.name" class="deploy_user-headimg">
            </div>
          </el-popover>
        </template>
      </el-table-column>
      <el-table-column
          prop="status"
          label="状态">
      </el-table-column>
      <el-table-column
          prop="created_at"
          label="开始时间">
      </el-table-column>
      <el-table-column
          prop="ended_at"
          label="结束时间">
      </el-table-column>
      <el-table-column
          label="操作"
          :render-header="tableAction"
          width="120">
        <template slot-scope="scope">
          <el-button @click="showRecord(scope.row)" type="success" icon="el-icon-info" size="small" circle></el-button>
        </template>
      </el-table-column>
    </el-table>
    <Paginate
        :api="fetchAllFunc"
        :params="{server_id: server_id}"
        @val-change="(data) => recordsData = data"
        >
    </Paginate>
    <RecordShow
        :title="recordShowTitle"
        :dialogFormVisible="dialogFormVisible"
        :record="currentRecord"
        v-if="currentRecord"
        @closeDialogForm="dialogFormVisible = false"
    >
    </RecordShow>
  </div>
</template>

<script>
  import ToolBar from '~/components/ToolBar/ToolBar.vue';
  import HelpHint from '~/components/HelpHint/HelpHint.vue';
  import RecordShow from './Show.vue';
  import Paginate from '~/components/Pagination/Paginate.vue';
  import DCApi from '~/api/deploy_record.js';

  export default {
    props: {
      server_id: String
    },
    data() {
      return {
        recordShowTitle: '详情',
        currentRecord: null,
        dialogFormVisible: false,
        recordsData: []
      }
    },
    created() {
      this.fetchAllFunc = DCApi.fetchAll
    },
    methods: {
      tableAction() {
        return this.$createElement('HelpHint', {
          props: {
            content: '详情'
          }
        }, '操作');
      },
      showRecord(data) {
        this.currentRecord = data;
        this.dialogFormVisible = true;
      }
    },
    components: {
      ToolBar, HelpHint, RecordShow, Paginate
    }
  }
</script>
<style lang="less">
  .deploy_user-headimg {
    width: 30px;
    height: 30px;
    border-radius: 120px;
  }
</style>
