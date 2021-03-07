
<template>
    <div>
        <p class="mt-5 text-secondary">AlarmDetail</p>
        <div v-if="alarm" class="mt-2">
                {{alarm.name}}
                {{alarm.hour}}
                {{alarm.min}}
                {{alarm.days}}
                {{alarm.on}}
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
            alarm : {}
            // alarms : alarmsList ,
            // adds : ['https://www.braveterry.com/wp-content/uploads/2015/12/vue.js.png', 'https://cdn-images-1.medium.com/max/1200/0*iBvb3FQRnC4Xdyv4.jpg', 'https://www.braveterry.com/wp-content/uploads/2015/12/vue.js.png']
        }
    },

    created()
	{
        console.log('entered alarm detail')
        console.log(this.$route.params)
		g.env().then(env => this.env = env)
        g.call("get_alarm",this.$route.params).then(
            function(alarm){
                console.log("INSIDE THEN ALARM")
                console.log(alarm)
               return alarm
            } ).then(alarm => this.alarm = alarm)
	},
}
</script>