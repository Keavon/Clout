var vue;

var clout = {
	el: "body",
	data: {
		player: {
			name: "",
			status: "connected"
		}
	},
	methods: {
		login: function() {
			
		}
	}
};

document.addEventListener("DOMContentLoaded", function() {
	vue = new Vue(clout);
});
