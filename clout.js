var vue;
var clout = {
	el: "body",
	data: {
		username: ""
	},
	methods: {
		login: function() {
			alert("Welcome, President " + this.username + "!");
		}
	}
};

document.addEventListener("DOMContentLoaded", function() {
	vue = new Vue(clout);
});
