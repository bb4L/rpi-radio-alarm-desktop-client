<template>
  <div class="card">
    <div class="card-body">
      <p class="card-title test-primary">
        AlarmDetail {{ this.$route.params["idx"] }}
      </p>
      <div v-if="alarm" class="container">
        <div class="row mb-2">
          <div class="col">
            <p>Name:</p>
          </div>
          <input
            class="form-control col"
            v-model="alarm.name"
            type="text"
            placeholder="name"
          />
        </div>

        <div class="row mb-2">
          <div class="col">
            <p>Time:</p>
          </div>
          <div class="col">
            <div class="row">
              <input
                class="form-control col"
                v-model="alarm.hour"
                type="number"
                placeholder="hour"
              />
              <div class="col text-center">:</div>
              <input
                class="form-control col"
                v-model="alarm.min"
                type="number"
                placeholder="minute"
              />
            </div>
          </div>
        </div>

        <div class="row mb-2">
          <div class="col">
            <p>Days:</p>
          </div>
          <input
            class="form-control col text-right"
            v-model="alarm.days"
            type="text"
            placeholder="days eg. '1,2,3,4'"
          />
        </div>

        <div class="row mb-2">
          <div class="col">
            <p>Active:</p>
          </div>
          <div class="col text-right">
            <div class="custom-control custom-switch">
              <input
                id="alarmOn"
                type="checkbox"
                class="custom-control-input"
                v-model="alarm.on"
                :disabled="disabled"
              />
              <label class="custom-control-label" v-bind:for="'alarmOn'">
                <div v-if="alarm.on">on</div>
                <div v-if="!alarm.on">off</div>
              </label>
            </div>
          </div>
        </div>
        <div class="row">
          <div class="col-6"></div>
          <div class="col" v-if="idx == -1">
            <div class="row">
              <button class="col btn btn-outline-success mr-2" @click="createAlarm()">
                Create Alarm
              </button>
            </div>
          </div>

          <div class="col" v-if="idx != -1">
            <div class="row">
              <button class="col btn btn-outline-success mr-2" @click="saveAlarm()">
                Save
              </button>
              <b-button v-b-modal.delete-modal class="col btn btn-outline-danger">Delete</b-button>
              <b-modal
                id="delete-modal"
                title="Delete Alarm"
                hide-header-close="true"
              >
                Do you really want to delete the account?
                <template #modal-footer>
                  <b-button
                    class="btn btn-outline-success"
                    block
                    @click="$bvModal.hide('delete-modal')"
                  >
                    Cancel
                  </b-button>
                  <b-button class="btn btn-danger" block @click="deleteAlarm()">
                    Delete
                  </b-button>
                </template>
              </b-modal>
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
      alarm: {},
      idx: -1,
    };
  },

  created() {
    let idx = this.$route.params["idx"];

    if (idx === parseInt(idx, 10).toString()) {
      this.idx = parseInt(idx, 10);
    }

    if (this.idx >= 0) {
      g.call("get_alarm", { idx: this.idx.toString() })
        .then((alarm) => {
          this.handleAlarm(alarm, "Alarm retrieved");
        })
        .catch(() => {
          this.helpers.errorToast(this, "Alarm", "could not get alarm");
        });
    } else {
      this.idx = -1;
      this.disabled = false;
    }
  },

  methods: {
    createAlarm() {
      this.disabled = true;
      this.handleAlarm();

      g.call("create_alarm", { alarm: JSON.stringify(this.alarm) })
        .then(() => {
          this.$router.push("/alarms");
          this.helpers.successToast(this, "Alarm", "Alarm created");
        })
        .catch(() => {
          this.disabled = false;
          this.helpers.errorToast(this, "Alarm", "Alarm could not be created");
        });
    },

    saveAlarm() {
      this.prepareAlarm();

      g.call("change_alarm", {
        alarm: JSON.stringify(this.alarm),
        idx: this.idx.toString(),
      })
        .then((alarm) => {
          this.handleAlarm(alarm, "Alarm saved");
        })
        .catch(() => {
          this.helpers.errorToast(this, "Alarm", "could not save alarm");
        });
    },

    deleteAlarm() {
      this.disabled = true;

      g.call("delete_alarm", { idx: this.idx })
        .then(() => {
          this.$router.push("/alarms");
          this.helpers.successToast(this, "Alarm", "Alarm deleted");
        })
        .catch(() => {
          this.$router.push("/alarms");
          this.helpers.errorToast(this, "Alarm", "Alarm could not be deleted");
          this.disabled = false;
        });
    },

    handleAlarm(alarm, msg) {
      this.alarm = JSON.parse(alarm);
      this.helpers.successToast(this, "Alarm", msg);
      this.helpers.delay(1000).then(() => (this.disabled = false));
    },

    prepareAlarm() {
      this.alarm.hour = parseInt(this.alarm.hour, 10);
      this.alarm.min = parseInt(this.alarm.min, 10);
      if (
        typeof this.alarm.days === "string" ||
        this.alarm.days instanceof String
      ) {
        this.alarm.days = this.alarm.days.split(",").map((n) => {
          return parseInt(n, 10);
        });
      }
    },
  },
};
</script>