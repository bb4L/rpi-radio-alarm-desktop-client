
<template>
    <div class="card-body">
    <p class="card-title text-primary">Alarms</p>
    <div class="card-text">
      <div class="container">
      <div v-for="(alarm, idx) in alarms" v-bind:key="idx">
        <div class="row m-2 align-items-center">
          <div class="col card border-primary">
            <div class="row align-items-center m-2">
                <div class="col-6 p-2">
                  <router-link  v-bind:to="'alarm/' + idx" class="list-group-item list-group-item-action" :is="disabled ? 'span' : 'router-link'">
                    {{ alarm.name }}
                    Time: {{ alarm.hour }}:{{ alarm.min }}

                    Days:{{ alarm.days }}
                  </router-link>
                </div>

                <div class="col-4 p-2">
                </div>

                <div class="col-2 justify-content-end">
                  <div class="custom-control custom-switch">
                    <input
                      type="checkbox"
                      class="custom-control-input"
                      v-bind:id="'alarmOn'+idx"
                      v-bind:idx=idx
                      v-on:change="switchAlarm"
                      v-model="alarm.on"
                      :disabled="disabled"
                    />
                    <label class="custom-control-label" v-bind:for="'alarmOn'+idx">
                      <div v-if="alarm.on">
                        on
                      </div>
                      <div v-if="!alarm.on">
                        off              
                      </div>
                    </label>
                  </div>
                </div>
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
    return {
      alarms: [],
      disabled: true,
    };
  },

  created() {
    g.env().then((env) => (this.env = env));
    g.call("get_alarms")
      .then((alarms) => {
        this.alarms = JSON.parse(alarms);
        this.successToast();
      })
      .catch(() => {});
  },

  methods: {
    successToast() {
      this.helpers.successToast(this, "Alarms", `Updated alarms`);
      const delay = (t) => new Promise((resolve) => setTimeout(resolve, t));
      delay(1000).then(() => {
        this.disabled = false;
      });
    },

    errorToast() {
      this.helpers.errorToast(this, "Alarms", "Could not update alarms");
    },

    switchAlarm: function (event) {
      console.log("switch alarm");
      // console.log(event);
      // console.log(event.srcElement.attributes);
      // console.log(event.srcElement.attributes.idx);
      // console.log(Number(event.srcElement.attributes.idx.value));
      // console.log(this.alarms[Number(event.srcElement.attributes.idx.value)]);
      let alarm = this.alarms[Number(event.srcElement.attributes.idx.value)];
      // console.log("alarm");
      console.log(alarm);
      // console.log(alarm.on);
      // this.alarms[Number(event.srcElement.attributes.idx.value)] = alarm;
      g.call("change_alarm", {
        alarm: JSON.stringify(alarm),
        idx: event.srcElement.attributes.idx.value,
      }).then((data) => {
        // JSON.parse(atob(alarm))
        console.log("data");
        console.log(data);
      });

      this.$emit("change", event.target.checked);
    },
  },
};
</script>