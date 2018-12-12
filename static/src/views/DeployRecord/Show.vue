<template>
  <div class="role-dialog">
    <el-dialog :title="title" :visible="dialogFormVisible" :show-close="false">
      <div class="record-attr-item">
        <strong style="margin-right: 18px">部署者:</strong>
        <a :href="record.deploy_user.github_url">
          <img :src="record.deploy_user.avatar_url" class="deploy_user-headimg">
        </a>
      </div>
      <div class="record-attr-item">
        <strong style="margin-right: 18px">状态:</strong> {{record.status}}
      </div>
      <div class="record-attr-item">
        <strong style="margin-right: 18px">开始时间:</strong>
        {{record.created_at}}
      </div>
      <div class="record-attr-item">
        <strong style="margin-right: 18px">结束时间:</strong>
        {{record.ended_at}}
      </div>
      <div class="record-attr-item" v-if="record.ended_at">
        <strong style="margin-right: 18px">耗时（s）:</strong>
        {{duration}}
      </div>
      <div class="record-attr-item">
        <strong style="margin-right: 18px">Compare:</strong>
        <a :href="record.compare">{{record.compare}}</a>
      </div>
      <el-collapse v-model="activeLog" accordion>
        <el-collapse-item title="Stdout" name="1">
          <div v-html="stdout" class="cmd-log"></div>
        </el-collapse-item>
        <el-collapse-item title="Stderr" name="2">
          <div v-html="stderr" class="cmd-log"></div>
        </el-collapse-item>
      </el-collapse>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" :loading="btnLoading" @click="closeDialogForm">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
  export default {
    name: 'ShowRecord',
    props: {
      dialogFormVisible: Boolean,
      record: Object,
      title: String,
    },
    data() {
      return {
        btnLoading: false,
        activeLog: 1,
      }
    },
    computed: {
      duration() {
        if (!(this.record.created_at && this.record.ended_at)) {
          return
        }

        let time1 = Date.parse(this.record.created_at)
        let time2 = Date.parse(this.record.ended_at)
        return (time2 - time1) / 1000
      },
      stdout() {
        return this.record.stdout.replace(/\r/gi, '<br>')
      },
      stderr() {
        return this.record.stderr.replace(/\r/gi, '<br>')
      }
    },
    methods: {
      closeDialogForm() {
        this.$emit("closeDialogForm")
      },
    },
    components: {}
  }
</script>
<style lang="less">
  .record-attr-item {
    margin-bottom: 10px;
  }
  .cmd-log {
    background: #212121;
    color: #edede3;
    padding: 10px;
    margin: 0 0 10.5px;
    font-size: 13px;
    line-height: 1.5;
    word-break: break-all;
  }
</style>
