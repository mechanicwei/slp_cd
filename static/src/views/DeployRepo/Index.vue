<template>
  <div class="role-list">
    <ToolBar>
      <el-button type="primary" icon="el-icon-plus" size="small" @click="editDeployRepo(false)">添加</el-button>
    </ToolBar>
    <el-table
        :data="deployReposData"
        border
        ref="table"
        style="width: 100%">
      <el-table-column
          prop="name"
          label="名称">
      </el-table-column>
      <el-table-column
          prop="github_url"
          label="Github Url">
      </el-table-column>
      <el-table-column
          prop="webhook_secret"
          label="Webhook Secret">
      </el-table-column>
      <el-table-column
          prop="openids"
          label="Openids">
      </el-table-column>
      <el-table-column
          label="操作"
          :render-header="tableAction"
          width="180">
        <template slot-scope="scope">
          <el-button @click="editDeployRepo(scope.row)" type="primary" icon="el-icon-edit" size="small" circle></el-button>
          <el-button @click="deleteRepo(scope.row.id)" type="danger" icon="el-icon-delete" size="small" circle></el-button>
          <el-button @click="$router.push(`deploy_repos/${scope.row.id}/servers`)" type="success" icon="el-icon-arrow-right" size="small" circle></el-button>
        </template>
      </el-table-column>
    </el-table>
    <DeployRpoEdit
        :title="roleEditTitle"
        :dialogFormVisible="dialogFormVisible"
        :data="currentEditRepo"
        @val-change="deployRepoEditChange"
        @cancel="dialogFormVisible = false"
    >
    </DeployRpoEdit>
  </div>
</template>

<script>
  import ToolBar from '~/components/ToolBar/ToolBar.vue';
  import HelpHint from '~/components/HelpHint/HelpHint.vue';
  import DeployRpoEdit from './Edit.vue';
  import DRApi from '~/api/deploy_repo.js';
  import Vue from 'vue';

  export default {
    data() {
      return {
        roleEditTitle: '编辑',
        currentEditRepo: false,
        dialogFormVisible: false,
        deployReposData: []
      }
    },
    created() {
      this.fetchData()
    },
    methods: {
      fetchData() {
        let self = this;
        DRApi.fetchAll((resp) => {
          self.deployReposData = resp.data
        })
      },
      tableAction() {
        return this.$createElement('HelpHint', {
          props: {
            content: '编辑 / 删除'
          }
        }, '操作');
      },
      deployRepoEditChange(data) {
        if (data.id) {
          this.updateDeployRpo(data)
        } else {
          this.createDeployRpo(data)
        }
        this.dialogFormVisible = false;
      },
      createDeployRpo(data) {
        let self = this;
        DRApi.create(data, (resp) => {
          self.$notify({
            title: '创建成功！',
            type: 'success'
          });
          self.deployReposData.push(resp.data)
          self.currentEditRepo = false;
        });
      },
      updateDeployRpo(data) {
        let self = this;
        DRApi.update(data, (resp) => {
          self.$notify({
            title: '更新成功！',
            type: 'success'
          });

          let index = self.deployReposData.findIndex((item) => item.id === resp.id)
          Vue.set(self.deployReposData, index, resp.data)
        });
      },
      editDeployRepo(data) {
        if (data) {
          this.currentEditRepo = data;
        } else {
          this.currentEditRepo = false;
        }

        this.dialogFormVisible = true;
      },
      deleteRepo(id) {
        let self = this;
        self.$confirm('确定删除?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          self.$notify({
            title: '成功',
            message: '假装被删除',
            type: 'success'
          });
          self.dialogFormVisible = false;
        }).catch(() => {

        });
      },
    },
    components: {
      ToolBar, HelpHint, DeployRpoEdit
    }
  }
</script>
<style lang="less">

</style>
