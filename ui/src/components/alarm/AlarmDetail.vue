
<template>
  <div class="card">
    <div class="card-body">
      <p class="card-title test-primary">
        AlarmDetail {{ this.$route.params["idx"] }}
      </p>
      <div v-if="alarm" class="container">
        <div class="row">
          <div class="col">
            <p>Name:</p>
          </div>
          <div class="col">
            {{ alarm.name }}
          </div>
        </div>
        <!-- TODO: make values editable -->
        <div class="row">
          <div class="col">
            <p>Time:</p>
          </div>
          <div class="col">{{ alarm.hour }}:{{ alarm.min }}</div>
        </div>

        <div class="row">
          <div class="col">
            <p>Days:</p>
          </div>
          <div class="col">
            {{ alarm.days }}
          </div>
        </div>

        <div class="row">
          <div class="col">
            <p>Active:</p>
          </div>
          <div class="col">
            {{ alarm.on }}
          </div>
        </div>

        <div class="row">
          <div class="col-10"></div>
          <div class="col">
            <button class="btn btn-outline btn-success" @click="saveAlarm()">
              Save
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import g from "guark";

export default {
  name: "Alarms",
  data() {
    // alarmsList = g.call("get_alarms",{})
    return {
      alarm: {},
      idx: -1,
    };
  },

  created() {
    console.log("entered alarm detail");
    console.log(this.$route.params);
    this.idx = this.params["idx"];
    g.env().then((env) => (this.env = env));
    g.call("get_alarm", this.$route.params)
      // .then(function (alarm) {
      //   console.log("INSIDE THEN ALARM");
      //   console.log(alarm);
      //   return alarm;
      // })
      .then((alarm) => {
        // btoa(JSON.stringify(alarm))
        // console.log()
        console.log("get alarm");
        console.log(alarm);
        this.alarm = JSON.parse(alarm);

        this.helpers.successToast(this, "Alarm", "retrieved alarm");
        console.log("before delay");
        console.log(this.alarm);
      });
  },

  methods: {
    saveAlarm() {
      g.call("change_alarm", {
        alarm: JSON.tringify(this.alarm),
        idx: this.idx,
      }).then((alarm) => {
        this.alarm = alarm;
        this.helper.successToast(this, "Alarm", "Alarm saved");
        this.helper.delay(1000).then(() => (this.disabled = false));
      });

      // this.$emit("click", event.target.checked);
    },
  },

  // TODO: delete alarm (with modal)
};
</script>