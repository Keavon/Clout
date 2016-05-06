var vue;

var clout = {
	el: "body",
	data: {
		player: {
			name: "",
			status: "stranger"
		}
	},
	methods: {
		login: function() {
			if (this.player.status === "connected") {
				return;
			}
			
			db = new Firebase("https://cloutgame.firebaseio.com/users/");
			
			db.once("value", function(ref) {
				if (ref.child(clout.data.player.name + "/country").exists()) {
					alert(clout.data.player.name + " is already in use");
					return;
				}
				
				db.child(clout.data.player.name + "/country").set(1);
				db.child(clout.data.player.name).onDisconnect().remove();
				
				clout.data.player.status = "connected";
			});
		}
	}
};

document.addEventListener("DOMContentLoaded", function() {
	vue = new Vue(clout);
});
