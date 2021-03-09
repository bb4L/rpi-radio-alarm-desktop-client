
<template>
  <div>
    <div class="card-body">
    <p class="card-title text-primary">Alarms</p>
    <div class="card-text">
        <div class="container">
      <div v-for="(alarm, idx) in alarms" v-bind:key="idx">
        <div class="row m-2">

          <div class="col card border-primary">
            <router-link  v-bind:to="'alarm/' + idx" class="list-group-item list-group-item-action">

            <div class="row align-items-center m-2">
                <div class="col">
              {{ alarm.name }}
              {{ alarm.hour }}
              {{ alarm.min }}
              {{ alarm.days }}
              {{ alarm.on }}
              {{ idx }}
              <!-- TODO: add button to switch alarm on/off -->
                </div>
                <button>test </button>
              <!-- <router-link
                class="btn btn-primary"
                v-bind:to="'alarm/' + idx"
                >alarm detail</router-link
              > -->
            </div>
          </router-link>

          </div>
          <div class="col">
              <button class="btn btn-outline-success">
                  test
              </button>
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
    // alarmsList = g.call("get_alarms",{})
    return {
      alarms: [],
      // alarms : alarmsList ,
      // adds : ['https://www.braveterry.com/wp-content/uploads/2015/12/vue.js.png', 'https://cdn-images-1.medium.com/max/1200/0*iBvb3FQRnC4Xdyv4.jpg', 'https://www.braveterry.com/wp-content/uploads/2015/12/vue.js.png']
    };
  },

  created() {
    console.log("entered");
    g.env().then((env) => (this.env = env));
    g.call("get_vals", {}).then(function (value) {
      console.log("got vals");
      console.log(value);
    });
    g.call("get_alarms")
      .then(function (alarms) {
        console.log("INSIDE THEN");
        console.log(alarms);
        return alarms;
      })
      .then((alarms) => (this.alarms = alarms));
  },
};
</script>