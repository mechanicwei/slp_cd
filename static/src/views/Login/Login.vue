<template>
  <div class="login">
    <div class="login-form">
      <div class="login-header">
        <img src="../../assets/images/logo.png" height="100" alt="">
        <p>{{ $Config.siteName }}</p>
      </div>
      <el-input
          placeholder="请输入用户名"
          suffix-icon="fa fa-user"
          v-model="userNmae"
          style="margin-bottom: 18px"
      >
      </el-input>

      <el-input
          placeholder="请输入密码"
          suffix-icon="fa fa-keyboard-o"
          v-model="password"
          type="password"
          style="margin-bottom: 18px"
          @keyup.native.enter="login"
      >
      </el-input>

      <el-button
          type="primary" :loading="loginLoading"
          style="width: 100%;margin-bottom: 18px"
          @click.native="login"
      >登录
      </el-button>
    </div>
  </div>
</template>

<script>
  import SessionApi from '~/api/session.js';

  export default {
    data() {
      return {
        userNmae: '',
        password: '',
        Remenber: true,
        loginLoading: false
      }
    },
    methods: {
      login() {
        let APP = this;
        APP.loginLoading = true;

        SessionApi.login({username: this.userNmae, password: this.password}, (resp) => {
          sessionStorage.setItem(APP.$Config.tokenKey, resp.data.token);
          APP.$notify({
            title: '登录成功',
            type: 'success',
          });
          APP.loginLoading = false;
          APP.$router.push({path: '/'});
        })
      }
    }
  }
</script>

<style lang="less">
  @import "Login.less";
</style>
