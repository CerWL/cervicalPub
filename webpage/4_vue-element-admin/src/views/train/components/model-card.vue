<template>
  <div class="card">
    <el-card class="box-card" shadow="hover">
      <div slot="header" class="clearfix">
        <span>Cell20198_231125978646576151</span>
      </div>
      <div class="flex model-info">
        <section class="info flex created">
          <h3>Created</h3>
          <b>2019.8.20</b>
          <i>1 computer hour</i>
        </section>
        <section class="info flex analyzed">
          <h3>Analyzed</h3>
          <b>5238 images</b>
          <i>2 labels,520 test images</i>
        </section>
        <section class="info flex avg">
          <h3>Avg precision</h3>
          <b>0.996</b>
        </section>
        <section class="percent">
          <el-progress type="circle" class="percent" :percentage="percentage1"></el-progress>
          <h4 class="percent-title">Precision</h4>
          <el-tooltip placement="bottom" class="percent-tip">
            <div slot="content">
              percent
              <br />表示精确度
            </div>
            <i class="el-icon-question"></i>
          </el-tooltip>
        </section>
        <section class="recall">
          <el-progress type="circle" class="recall" :percentage="percentage2"></el-progress>
          <h4 class="recall-title">Recall</h4>
          <el-tooltip placement="bottom" class="recall-tip">
            <div slot="content">
              recall
              <br />表示有点东西
            </div>
            <i class="el-icon-question"></i>
          </el-tooltip>
        </section>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "card",
  components: {},
  data() {
    return {
      percentage1: 10,
      percentage2: 10
    }
  },
  methods: {},
  mounted() {
    setTimeout(() => { this.percentage1 = 27;this.percentage2 = 20 }, 4000)
    setTimeout(() => { this.percentage1 = 64;this.percentage2 = 40 }, 4500)
    setTimeout(() => { this.percentage1 = 88;this.percentage2 = 66 }, 5000)
    setTimeout(() => { this.percentage1 = 95;this.percentage2 = 96 }, 5500)
  }
}
</script>

<style lang="scss" scoped>
.card {
  i {
    color: #ccc;
    font-style: normal;
  }
  b {
    font-size: 22px;
  }
  .model-info {
    justify-content: space-around;
    .info {
      flex-direction: column;
      justify-content: flex-start;
    }
    .percent {
      position: relative;
    }
    .percent-title {
      position: absolute;
      top: 10px;
      left: 30px;
    }
    .el-icon-question {
      position: absolute;
      bottom: 30px;
      left: 55px;
    }
    .recall {
      position: relative;
    }
    .recall-title {
      position: absolute;
      top: 10px;
      left: 40px;
    }
  }
}
</style>
