
<template>
    <div>
        <p class="mt-5 text-secondary">Alarms</p>
        <div class="mt-2">
            <div v-for="alarm in alarms">
                {{alarm.name}}
                {{alarm.hour}}
                {{alarm.min}}
                {{alarm.days}}
                {{alarm.on}}
                <!-- <a href="#"> -->
                    <!-- <p v-bind:src="alarm" style="width: 150px;" class="d-block mb-3">
                        {{alarm.name}}
                    </p> -->
                <!-- </a> -->
            </div>
        </div>
    </div>
</template>

<script>
import g from 'guark'

export default {
    name: "Alarms",
    data(){
        // alarmsList = g.call("get_alarms",{})
        return {
            alarms : []
            // alarms : alarmsList ,
            // adds : ['https://www.braveterry.com/wp-content/uploads/2015/12/vue.js.png', 'https://cdn-images-1.medium.com/max/1200/0*iBvb3FQRnC4Xdyv4.jpg', 'https://www.braveterry.com/wp-content/uploads/2015/12/vue.js.png']
        }
    },

    created()
	{
        console.log('entered')
		g.env().then(env => this.env = env)
        g.call("get_vals",{}).then(
            function(value){
                console.log("got vals")
                console.log(value)
            }
        )
        g.call("get_alarms",{}).then(
            function(alarms){
                console.log("INSIDE THEN")
                console.log(alarms)
               return alarms
            } ).then(alarms => this.alarms = alarms)
	},
}
</script>