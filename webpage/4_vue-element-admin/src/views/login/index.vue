<template>
  <div class="login-container">
    <loginHeader></loginHeader>
    <section class="main flex">
      <img class="img" src="../../assets/login.png" />
      <el-form
        ref="loginForm"
        :model="loginForm"
        :rules="loginRules"
        class="login-form"
        autocomplete="on"
        label-position="left"
      >
        <div class="title-container">
          <h3 class="title">讯动医疗</h3>
        </div>

        <el-form-item prop="username">
          <span class="svg-container">
            <svg-icon icon-class="user" />
          </span>
          <el-input
            ref="username"
            v-model="loginForm.username"
            placeholder="请输入邮箱"
            name="username"
            type="email"
            tabindex="1"
            autocomplete="on"
          />
        </el-form-item>

        <el-form-item prop="password">
          <span class="svg-container">
            <svg-icon icon-class="password" />
          </span>
          <el-input
            :key="passwordType"
            ref="password"
            v-model="loginForm.password"
            :type="passwordType"
            placeholder="密码"
            name="password"
            tabindex="2"
            autocomplete="on"
            @keyup.enter.native="handleLogin"
          />
          <span class="show-pwd" @click="showPwd">
            <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
          </span>
        </el-form-item>

        <el-form-item prop="password" v-if="register">
          <span class="svg-container">
            <svg-icon icon-class="password" />
          </span>
          <el-input
            :key="passwordType"
            ref="password"
            v-model="loginForm.confirmPassword"
            :type="passwordType"
            placeholder="确认密码"
            name="password"
            tabindex="2"
            autocomplete="on"
            @keyup.enter.native="handleLogin"
          />
          <span class="show-pwd" @click="showPwd">
            <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
          </span>
        </el-form-item>

        <el-form-item prop="code" v-if="register">
          <span class="svg-container">
            <svg-icon icon-class="email" />
          </span>
          <el-input
            key="code"
            ref="code"
            maxlength="6"
            v-model="loginForm.code"
            type="text"
            placeholder="邮箱验证码"
            name="code"
            tabindex="2"
            autocomplete="on"
          />
          <span class="show-pwd">
            <svg-icon icon-class="guide" @click="sendCode" v-if="time === 60" />
            <i v-else>{{time}}</i>
          </span>
        </el-form-item>

        <el-button
          :loading="loading"
          type="primary"
          style="width:100%;margin-bottom:20px;"
          @click.native.prevent="handleLogin"
        >{{register?'注册并登录':'登录'}}</el-button>
        <div style="position:relative">
          <div class="tips">
            <el-link type="primary" :underline="false" icon="el-icon-key">忘记密码</el-link>
            <el-link
              type="primary"
              :underline="false"
              icon="el-icon-user"
              class="register"
              @click="handleRegisterLogin"
            >{{!register?'注册':'登录'}}</el-link>
          </div>
        </div>
      </el-form>
    </section>
    <loginFooter></loginFooter>
  </div>
</template>

<script>
import { validEmail } from '@/utils/validate'
import loginHeader from './components/login-header'
import loginFooter from './components/login-footer'
import { getCode } from '@/api/user'
let timer

export default {
  name: 'Login',
  components: { loginHeader, loginFooter },
  data() {
    const validateUsername = (rule, value, callback) => {
      if (validEmail(value) || this.loginForm.username === 'admin') {
        callback()
      } else {
        callback(new Error('请输入正确的邮箱号'))
      }
    }
    const validatePassword = (rule, value, callback) => {
      if (value.length < 3) {
        callback(new Error('至少3位'))
      } else {
        callback()
      }
    }
    const validateConfirmPassword = (rule, value, callback) => {
      if (this.loginForm.password === this.loginForm.confirmPassword) {
        callback(new Error('两次输入的密码不同'))
      } else {
        callback()
      }
    }
    return {
      fuValue: '',
      register: false,
      loginForm: {
        username: '',
        code: '',
        password: '',
        confirmPassword: ''
      },
      time: 60,
      loginRules: {
        username: [
          { required: true, trigger: 'blur', validator: validateUsername }
        ],
        password: [
          { required: true, trigger: 'blur', validator: validatePassword }
        ],
        confirmPassword: [
          { required: true, trigger: 'blur', validator: validateConfirmPassword }
        ]
      },
      passwordType: 'password',
      capsTooltip: false,
      loading: false,
      showDialog: false,
      redirect: undefined,
      otherQuery: {}
    }
  },
  watch: {
    $route: {
      handler(route) {
        const query = route.query
        if (query) {
          this.redirect = query.redirect
          this.otherQuery = this.getOtherQuery(query)
        }
      },
      immediate: true
    }
  },
  created() {
    // window.addEventListener('storage', this.afterQRScan)
  },
  mounted() {
    if (this.loginForm.username === '') {
      this.$refs.username.focus()
    } else if (this.loginForm.password === '') {
      this.$refs.password.focus()
    }
  },
  destroyed() {
    clearInterval(timer)
    // window.removeEventListener('storage', this.afterQRScan)
  },
  methods: {
    handleRegisterLogin() {
      this.register = !this.register
      this.loginForm.username = ''
      this.loginForm.password = ''
    },
    sendCode() {
      if (!validEmail(this.loginForm.username)) return
      getCode({ 'email': this.loginForm.username })
      timer = setInterval(() => {
        this.time--
        if (this.time === 0) {
          clearInterval(timer)
          this.time = 60
        }
      }, 1e3)
    },
    showPwd() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
      this.$nextTick(() => {
        this.$refs.password.focus()
      })
    },
    handleLogin() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.loading = true
          if (this.register) {
            this.$store
              .dispatch('user/register', this.loginForm)
              .then(() => {
                setTimeout(() => {
                  this.$store
                    .dispatch('user/login', this.loginForm)
                    .then(() => {
                      this.$router.push({
                        path: this.redirect || '/',
                        query: this.otherQuery
                      })
                      this.loading = false
                    })
                    .catch(() => {
                      this.loading = false
                    })
                }, 1e3)
              })
              .catch(() => {
                this.loading = false
              })
          } else {
            this.$store
              .dispatch('user/login', this.loginForm)
              .then(() => {
                this.$router.push({
                  path: this.redirect || '/',
                  query: this.otherQuery
                })
                this.loading = false
              })
              .catch(() => {
                this.loading = false
              })
          }
        } else {
          return false
        }
      })
    },
    getOtherQuery(query) {
      return Object.keys(query).reduce((acc, cur) => {
        if (cur !== 'redirect') {
          acc[cur] = query[cur]
        }
        return acc
      }, {})
    }
    // afterQRScan() {
    //   if (e.key === 'x-admin-oauth-code') {
    //     const code = getQueryObject(e.newValue)
    //     const codeMap = {
    //       wechat: 'code',
    //       tencent: 'code'
    //     }
    //     const type = codeMap[this.auth_type]
    //     const codeName = code[type]
    //     if (codeName) {
    //       this.$store.dispatch('LoginByThirdparty', codeName).then(() => {
    //         this.$router.push({ path: this.redirect || '/' })
    //       })
    //     } else {
    //       alert('第三方登录失败')
    //     }
    //   }
    // }
  }
}
</script>

<style lang="scss">
/* 修复input 背景不协调 和光标变色 */
/* Detail see https://github.com/PanJiaChen/vue-element-admin/pull/927 */

$bg: #5b646e;
$light_gray: #fff;
$cursor: #fff;

@supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
  .login-container .el-input input {
    color: $cursor;
  }
}
.register {
  margin-left: 20px;
}
/* reset element-ui css */
.login-container {
  .main {
    justify-content: space-around;
    .img {
      width: 36vw;
      margin-top: 150px;
    }
  }
  .el-input {
    display: inline-block;
    height: 47px;
    width: 85%;

    input {
      background: transparent;
      border: 0px;
      -webkit-appearance: none;
      border-radius: 0px;
      padding: 12px 5px 12px 15px;
      color: $light_gray;
      height: 47px;
      caret-color: $cursor;

      &:-webkit-autofill {
        box-shadow: 0 0 0px 1000px $bg inset !important;
        -webkit-text-fill-color: $cursor !important;
      }
    }
  }

  .el-form-item {
    border: 1px solid rgba(255, 255, 255, 0.1);
    background: rgba(0, 0, 0, 0.1);
    border-radius: 5px;
    color: #454545;
  }
}
</style>

<style lang="scss" scoped>
$bg: #2f3135;
$dark_gray: #889aa4;
$light_gray: #eee;

.login-container {
  min-height: 100%;
  width: 100%;
  background-color: $bg;
  overflow: hidden;

  .login-form {
    position: relative;
    width: 400px;
    max-width: 100%;
    padding: 160px 35px 0;
    overflow: hidden;
  }

  .tips {
    font-size: 14px;
    color: #fff;
    margin-bottom: 10px;

    span {
      &:first-of-type {
        margin-right: 16px;
      }
    }
  }

  .svg-container {
    padding: 6px 5px 6px 15px;
    color: $dark_gray;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
  }

  .title-container {
    position: relative;

    .title {
      font-size: 26px;
      color: $light_gray;
      margin: 0px auto 40px auto;
      text-align: center;
      font-weight: bold;
    }
  }

  .show-pwd {
    position: absolute;
    right: 10px;
    top: 7px;
    font-size: 16px;
    color: $dark_gray;
    cursor: pointer;
    user-select: none;
  }

  .thirdparty-button {
    position: absolute;
    right: 0;
    bottom: 6px;
  }

  @media only screen and (max-width: 470px) {
    .thirdparty-button {
      display: none;
    }
  }
}
</style>
