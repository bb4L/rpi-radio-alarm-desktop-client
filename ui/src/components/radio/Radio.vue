<template>
  <div>
    <div class="card-body">
      <p class="card-title text-primary">Radio</p>
      <div class="mt-2">
        <div v-if="radio" class="container">
          <div class="custom-control custom-switch">
            <input
              type="checkbox"
              class="custom-control-input"
              id="radioRunning"
              v-on:change="switchRadio"
              :disabled="disabled"
            />
            <label class="custom-control-label" for="radioRunning"
              >is playing
            </label>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import g from "guark";

export default {
  name: "Radio",
  data() {
    return {
      radio: {},
      disabled: true,
    };
  },

  created() {
    g.call("get_radio")
      .then((radio) => {
        this.radio = radio;
        this.successToast();
      })
      .catch(() => {
        this.errorToast();
      });
  },
  methods: {
    switchRadio: function (event) {
      this.disabled = true;

      if (event.srcElement.checked) {
        g.call("start_radio")
          .then((radio) => {
            this.radio = radio;
            this.successToast();
          })
          .catch(() => {
            this.errorToast();
          });
      } else {
        g.call("stop_radio")
          .then((radio) => {
            this.radio = radio;
            this.successToast();
          })
          .catch(() => {
            this.errorToast();
          });
      }
      this.$emit("change", event.target.checked);
    },

    errorToast() {
      this.helpers.errorToast(
        this,
        "Radio",
        `Could not update radio information`
      );
    },

    successToast() {
      this.helpers.successToast(this, "Radio", `Updated radio information`);
      this.helpers.delay(2000).then(() => {
        this.disabled = false;
      });
    },
  },
};
</script>