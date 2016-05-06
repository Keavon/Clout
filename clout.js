var vue;

var clout = {
	el: "body",
	data: {
		username: "",
		playing: false
	},
	methods: {
		login: function() {
			if(this.playing) {
				alert("You're already playing!");
				return;
			}

			db = new Firebase("https://apesclout.firebaseio.com/users/");

			db.once("value", function(ref) {

				if(ref.child(clout.data.username + "/country").exists()) {
					alert(clout.data.username + " is already in use");
					return;
				}

				db.child(clout.data.username + "/country").set(1);
				alert("Welcome, President " + clout.data.username + "!");

				db.child(clout.data.username).onDisconnect().remove();

				clout.data.playing = true;
			});
		}
	}
};

document.addEventListener("DOMContentLoaded", function() {
	vue = new Vue(clout);
});
