<template>
  <div class="role-list">
    <ToolBar>
      <el-button type="primary" icon="el-icon-plus" size="small" @click="editDeployServer(false)">添加</el-button>
    </ToolBar>
    <el-table
        :data="deployServersData"
        border
        ref="table"
        style="width: 100%">
      <el-table-column
          prop="name"
          label="名称">
      </el-table-column>
      <el-table-column
          prop="branch"
          label="分支">
      </el-table-column>
      <el-table-column
          prop="deploys_count"
          label="部署次数">
      </el-table-column>
      <el-table-column
          label="操作"
          :render-header="tableAction"
          width="180">
        <template slot-scope="scope">
          <el-button @click="editDeployServer(scope.row)" type="primary" icon="el-icon-edit" size="small" circle></el-button>
          <el-button @click="deleteServer(scope.row.id)" type="danger" icon="el-icon-delete" size="small" circle></el-button>
          <el-button @click="$router.push(`/deploy_servers/${scope.row.id}/records`)" type="success" icon="el-icon-arrow-right" size="small" circle></el-button>
        </template>
      </el-table-column>
    </el-table>
    <DeployServerEdit
        :title="roleEditTitle"
        :dialogFormVisible="dialogFormVisible"
        :data="currentEditServer"
        @val-change="deployServerEditChange"
        @cancel="dialogFormVisible = false"
    >
    </DeployServerEdit>
  </div>
</template>

<script>
  import ToolBar from '~/components/ToolBar/ToolBar.vue';
  import HelpHint from '~/components/HelpHint/HelpHint.vue';
  import DeployServerEdit from './Edit.vue';
  import DSApi from '~/api/deploy_server.js';
  import Vue from 'vue';

  export default {
    props: {
      repo_id: String
    },
    data() {
      return {
        roleEditTitle: '编辑',
        currentEditServer: false,
        dialogFormVisible: false,
        deployServersData: []
      }
    },
    created() {
      this.fetchData()
    },
    methods: {
      fetchData() {
        let self = this;
        DSApi.fetchAll({repo_id: this.repo_id}, (resp) => {
          self.deployServersData = resp.data
        })
      },
      tableAction() {
        return this.$createElement('HelpHint', {
          props: {
            content: '编辑 / 删除'
          }
        }, '操作');
      },
      deployServerEditChange(data) {
        if (data.id) {
          this.updateDeployServer(data)
        } else {
          this.createDeployServer(data)
        }
        this.dialogFormVisible = false;
      },
      createDeployServer(data) {
        let self = this;
        let repo_id = parseInt(this.repo_id)
        DSApi.create(Object.assign({deploy_repo_id: repo_id}, data), (resp) => {
          self.$notify({
            title: '创建成功！',
            type: 'success'
          });
          self.deployServersData.push(resp.data)
          self.currentEditServer = false;
        });
      },
      updateDeployServer(data) {
        let self = this;
        DSApi.update(data, (resp) => {
          self.$notify({
            title: '更新成功！',
            type: 'success'
          });

          let index = self.deployServersData.findIndex((item) => item.id === resp.id)
          Vue.set(self.deployServersData, index, resp.data)
        });
      },
      editDeployServer(data) {
        if (data) {
          this.currentEditServer = data;
        } else {
          this.currentEditServer = false;
        }

        this.dialogFormVisible = true;
      },
      deleteServer(id) {
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
      ToolBar, HelpHint, DeployServerEdit
    }
  }
</script>
<style lang="less">

</style>
