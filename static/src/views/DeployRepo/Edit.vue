<template>
  <div class="role-dialog">
    <el-dialog :title="title" :visible="dialogFormVisible" :show-close="false" width="45%">
      <el-form :model="form">
        <el-form-item label="名称">
          <el-input v-model="form.name" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="Github Url">
          <el-input v-model="form.github_url" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="Webhook Secret">
          <el-input v-model="form.webhook_secret" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="Openids（用分号;隔开）">
          <el-input v-model="form.openids" auto-complete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="$emit('cancel');">取 消</el-button>
        <el-button type="primary" :loading="btnLoading" @click="btnOk">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
  export default {
    name: 'Edit',
    props: {
      dialogFormVisible: Boolean,
      data: [Object, Boolean],
      title: String,
    },
    data() {
      return {
        form: {
          id: null,
          name: '',
          github_url: '',
          webhook_secret: '',
          openids: '',
        },
        btnLoading: false,
      }
    },
    watch: {
      'data': {
        handler: function () {
          if (this.data) {
            for (let k in this.form) {
              this.form[k] = this.data[k];
            }
          } else {
            for (let k in this.form) {
              this.form[k] = '';
            }
          }
        },
        deep: true
      }
    },
    methods: {
      handleChange() {
      },
      btnOk() {
        this.$emit('val-change', this.form);
        this.btnLoading = true;
        setTimeout(() => {
          this.btnLoading = false
        }, 1000)
      },
    },
    components: {}
  }
</script>
<style lang="less">
  .role-dialog {

  }
</style>
