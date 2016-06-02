var vue;

var clout = {
	el: "body",
	data: {
		connection: {
			name: "",
			room: "",
			status: "stranger",
			token: "",
			creatingRoom: false,
			invitePrompt: false
		},
		country: {
			name: "",
			rank: 0,
			demand: 0,
			money: 0,
			damage: 0
		},
		resources: [
			{
				name: "Coal",
				capacity: 0,
				owned: 0,
				operating: 0,
				damage: 0,
				revenue: 0,
				cost: 0,
				icon: '<circle cx="500" cy="500" r="475" fill="#fff"/><path d="M783.7 546.1l43.8-132.2-112.3-121.6L490.9 205 246.3 330.4 150 486l40.8 98.5-5.7 84 72.5 56 150.6 21.3 95.1-29.1L638.1 795l156.4-75 55.5-83.7-66.3-90.2zM382.3 293l112.4-42 16.1 42H382.3zm52.4 181.6l91.8 12.3 7.4 42-99.2-54.3zM690 315.8l93.7 89-112-6.9-94.4-53.7-36.8-83.2L690 315.8zm92.7 374l-130.6-2.1L556.3 552l126.3 98.7 94.2-60.4 34.3 44.7-28.4 54.8z"/><path d="M500 0C223.9 0 0 223.9 0 500s223.9 500 500 500 500-223.9 500-500S776.1 0 500 0zm0 940C257 940 60 743 60 500S257 60 500 60s440 197 440 440-197 440-440 440z"/>'
			},
			{
				name: "Oil",
				capacity: 0,
				owned: 0,
				operating: 0,
				damage: 0,
				revenue: 0,
				cost: 0,
				icon: '<circle cx="500" cy="500" r="475" fill="#fff"/><ellipse cx="410.1" cy="206.7" rx="38.3" ry="14.3"/><path d="M713.2 495.2c9-9.3 14.3-19.3 14.3-29.8 0-10.4-5.3-20.5-14.3-29.7V264.9c9-9.2 14.2-19.9 14.2-29.9h-.2c0-48-101.8-85.1-227.4-85.1S272.4 187.2 272.4 235h.3c.2 10.1 5.1 19.9 14.1 28.9v171.9c-9 9.2-13.9 19.2-13.9 29.6 0 10.3 4.9 20.2 13.9 29.4v240.6c-9 9.2-14.1 19.2-14.1 29.6 0 47 101.7 85.1 227.2 85.1S727.4 812 727.4 765c0-10.4-5.3-20.5-14.3-29.7V495.2zm-22.9 282.9c-10.7 7.7-26 15-44.1 21.2-39.9 13.5-91.9 20.9-146.3 20.9-54.4 0-106.4-7.4-146.3-20.9-18.2-6.1-33.4-13.5-44.1-21.2-8.8-6.3-12-11.1-13.1-13.1.5-1 1.7-2.7 3.7-5 2.1 2.2 5.1 5 9.4 8.2 10.7 7.7 25.9 15 44.1 21.2 39.9 13.5 92 20.9 146.5 20.9 14.9 0 29-.6 44-1.6V548.9c44-3.2 83-11 113.1-22.1v258.4c13-5.2 24.6-11 33-17.1 4.4-3.2 7.5-6 9.5-8.2 2.1 2.3 3.2 4 3.7 5-1.1 2.1-4.3 6.8-13.1 13.2zm0-299.5c-10.7 7.7-26 15-44.1 21.2-39.9 13.5-91.9 20.9-146.3 20.9-54.4 0-106.4-7.4-146.3-20.9-18.2-6.1-33.4-13.5-44.1-21.2-8.8-6.3-12-11.1-13.1-13.1.5-1 1.7-2.7 3.7-5 2.1 2.2 5.1 5 9.4 8.2 10.7 7.7 25.9 15 44.1 21.2 39.9 13.5 92 20.9 146.5 20.9 14.9 0 29-.6 44-1.6V318.5c44-3.2 83-11 113.1-22v189.2c13-5.2 24.6-11 33-17.1 4.4-3.2 7.5-6 9.5-8.2 2.1 2.3 3.2 4 3.7 5-1.1 2-4.3 6.8-13.1 13.2zm0-230.5c-10.7 7.7-26 14.8-44.1 21-39.9 13.5-91.9 20.8-146.3 20.8-54.4 0-106.4-7.5-146.3-21-18.2-6.1-33.4-13.5-44.1-21.2-8.8-6.3-12-11.1-13.1-13.1 13.2-42.8 87-64.6 203.5-64.6 116.6 0 190.4 22.5 203.6 65.3-1.1 2-4.4 6.5-13.2 12.8z"/><path d="M500 0C223.9 0 0 223.9 0 500s223.9 500 500 500 500-223.9 500-500S776.1 0 500 0zm0 940C257 940 60 743 60 500S257 60 500 60s440 197 440 440-197 440-440 440z"/>'
			},
			{
				name: "Natural Gas",
				capacity: 0,
				owned: 0,
				operating: 0,
				damage: 0,
				revenue: 0,
				cost: 0,
				icon: '<circle cx="500" cy="500" r="475" fill="#fff"/><path d="M564.4 842.5C702.9 791.4 722.2 660.7 680.5 494 620.7 254.5 420 150 420 150s59.5 142.7 16.4 222.3C406 428.4 315.8 509 302.4 600.6 289 692.3 328 835 502.1 850c0 0-102.9-111.2-59.7-235.8S588.2 473 588.2 473s-52.1 57.1-32.8 102.1c20.3 47.3 38.5 55.5 53.3 98.1 30.2 86.7-44.3 169.3-44.3 169.3z"/><path d="M500 0C223.9 0 0 223.9 0 500s223.9 500 500 500 500-223.9 500-500S776.1 0 500 0zm0 940C257 940 60 743 60 500S257 60 500 60s440 197 440 440-197 440-440 440z"/>'
			},
			{
				name: "Nuclear",
				capacity: 0,
				owned: 0,
				operating: 0,
				damage: 0,
				revenue: 0,
				cost: 0,
				icon: '<circle cx="500" cy="500" r="475" fill="#fff"/><circle cx="500" cy="500" r="60"/><path d="M500 0C223.9 0 0 223.9 0 500s223.9 500 500 500 500-223.9 500-500S776.1 0 500 0zm0 940C257 940 60 743 60 500S257 60 500 60s440 197 440 440-197 440-440 440z"/><path d="M500 164c16.4 0 32.8 7.9 48.6 23.5 16.6 16.3 31.6 40.1 44.6 70.6 13.4 31.2 23.9 67.6 31.2 108.2 7.7 42.3 11.6 87.3 11.6 133.7s-3.9 91.4-11.6 133.7c-7.4 40.6-17.9 77.1-31.2 108.2-13.1 30.5-28.1 54.2-44.6 70.6-15.8 15.6-32.2 23.5-48.6 23.5s-32.8-7.9-48.6-23.5c-16.6-16.3-31.6-40.1-44.6-70.6-13.4-31.2-23.9-67.6-31.2-108.2-7.7-42.3-11.6-87.3-11.6-133.7s3.9-91.4 11.6-133.7c7.4-40.6 17.9-77.1 31.2-108.2 13.1-30.5 28.1-54.2 44.6-70.6 15.8-15.6 32.2-23.5 48.6-23.5m0-14c-82.8 0-150 156.7-150 350s67.2 350 150 350 150-156.7 150-350-67.2-350-150-350z"/><path d="M699.9 296c23.5 0 43.5 3.3 59.3 9.6 15.2 6.1 26.2 15 32.8 26.4 8.2 14.2 9.5 32.3 4 53.8-5.9 22.5-18.9 47.4-38.8 73.9-20.3 27.2-46.6 54.5-78.1 81.2-32.8 27.8-69.8 53.7-110 76.9-47.2 27.3-97 49.5-144 64.4-45 14.2-87.5 21.8-122.9 21.8-23.5 0-43.5-3.2-59.3-9.6-15.2-6.1-26.2-15-32.8-26.4-8.2-14.2-9.5-32.3-4-53.8 5.9-22.5 18.9-47.4 38.8-73.9 20.3-27.2 46.6-54.5 78.1-81.2 32.8-27.8 69.8-53.7 110-76.9 47.2-27.3 97-49.6 144-64.4 45-14.3 87.5-21.8 122.9-21.8h.1m-.1-14c-73.6 0-174.1 30.4-273.9 88.1C258.6 466.7 156.5 603.3 197.9 675c16.7 29 54.3 43 104.2 43 73.6 0 174.1-30.4 273.9-88.1C743.4 533.3 845.5 396.7 804.1 325c-16.7-29-54.3-43-104.2-43z"/><path d="M302.6 295c35.4 0 77.9 7.6 122.9 21.9 46.9 14.9 96.7 37.2 144 64.5 40.2 23.2 77.2 49.1 110 76.9 31.5 26.7 57.8 54 78.1 81.2 19.9 26.5 32.9 51.4 38.8 73.9 5.6 21.5 4.3 39.5-4 53.8-6.6 11.4-17.6 20.3-32.8 26.4-15.9 6.4-35.8 9.6-59.3 9.6-35.4 0-77.9-7.5-122.9-21.8-46.9-14.9-96.7-37.1-144-64.4-40.2-23.2-77.2-49.1-110-76.9-31.5-26.7-57.8-54-78.1-81.2-19.9-26.5-32.9-51.4-38.8-73.9-5.6-21.5-4.3-39.5 4-53.8 6.6-11.4 17.6-20.4 32.8-26.4 15.9-6.5 35.8-9.8 59.3-9.8h.4m-.4-13.8c-49.9 0-87.5 14-104.2 43C157 395.9 259.1 532.4 426.5 629c99.8 57.6 200.2 88.1 273.9 88.1 49.9 0 87.5-14 104.2-43 41.4-71.7-60.7-208.3-228.1-304.9-99.8-57.6-200.2-88-273.9-88z"/>'
			},
			{
				name: "Geothermal",
				capacity: 0,
				owned: 0,
				operating: 0,
				damage: 0,
				revenue: 0,
				cost: 0,
				icon: '<path fill="#fff" d="M500 50C251.5 50 50 251.5 50 500l165.1 410.9c11.4 7.9 23.2 15.4 35.3 22.4l162.3 59v.1c14.1 2.5 28.4 4.4 42.9 5.7l183.6-17.7c15-4.3 29.8-9.4 44.2-15.1L810 876l140-376c0-248.5-201.5-450-450-450z"/><path d="M500 0C223.9 0 0 223.9 0 500c0 170.3 85.1 320.6 215.1 410.9 12.4-14.2 28.6-27.3 48.3-38.7l14.9-8.6c26.6-19.1 41.1-41.8 41-65.1 0-27.6-20.3-54.4-57-75.5-50.1-28.8-77.6-68.3-77.6-111.2 0-42.6 27.3-81.8 76.7-110.4L246.8 476l98.3.1-49.1 85.1-13.9-24c-36.2 20.9-56.1 47.4-56.1 74.6 0 27.5 20.2 54.3 56.9 75.4 50.1 28.8 77.7 68.3 77.7 111.2.1 42.7-27.1 81.9-76.6 110.5l-15.1 8.7c-7 5-13.2 10.3-18.5 15.7 49.4 28.5 104.1 48.8 162.3 59-.2-2.8-.4-5.6-.4-8.4-.1-43.1 27.5-82.8 77.7-111.7l14.9-8.6c26.6-19.1 41.1-41.8 41-65.1 0-27.6-20.3-54.4-57-75.5-50.1-28.8-77.6-68.3-77.6-111.2 0-42.6 27.3-81.8 76.7-110.4L473.5 476l98.3.1-49.1 85.2-13.9-24c-36.2 20.9-56.1 47.4-56.1 74.6 0 27.5 20.2 54.3 56.9 75.4 50.1 28.8 77.7 68.3 77.7 111.2.1 42.7-27.1 81.9-76.6 110.5l-15.1 8.7c-27.2 19.5-42 42.5-41.9 66.2 0 4.8.6 9.5 1.9 14.2 14.6 1.3 29.4 1.9 44.4 1.9 48.3 0 95-6.9 139.2-19.6 1.5-41.7 28.9-79.9 77.6-108.1l14.9-8.6c26.6-19.1 41.1-41.8 41-65.1 0-27.6-20.3-54.4-57-75.5-50.1-28.8-77.6-68.3-77.6-111.2 0-42.6 27.3-81.8 76.7-110.4L700.2 476l98.3.1-49.1 85.2-13.9-24c-36.2 20.9-56.1 47.4-56.1 74.6 0 27.5 20.2 54.3 56.9 75.4C786.4 716 814 755.5 814 798.5c.1 42.7-27.1 81.9-76.6 110.5l-15.1 8.7c-20 14.3-33.3 30.6-38.9 47.6C868.8 892.1 1000 711.4 1000 500 1000 223.9 776.1 0 500 0zM98.8 319C167.8 166.3 321.5 60 500 60s332.2 106.3 401.2 259H98.8z"/>'
			},
			{
				name: "Solar",
				capacity: 0,
				owned: 0,
				operating: 0,
				damage: 0,
				revenue: 0,
				cost: 0,
				icon: '<circle cx="500" cy="500" r="475" fill="#fff"/><path d="M500 0C223.9 0 0 223.9 0 500s223.9 500 500 500 500-223.9 500-500S776.1 0 500 0zm0 940C257 940 60 743 60 500S257 60 500 60s440 197 440 440-197 440-440 440z"/><circle cx="500" cy="500" r="155.6"/><path d="M490 694h20v156h-20zM490 150h20v156h-20zM694 490h156v20H694zM150 490h156v20H150zM630.052 644.25l14.142-14.142 110.308 110.308-14.142 14.142zM245.485 259.584l14.142-14.142L369.934 355.75l-14.142 14.142zM630.1 355.8l110.31-110.31 14.14 14.143L644.245 369.94zM245.435 740.366L355.743 630.06l14.142 14.14L259.577 754.51z"/>'
			},
			{
				name: "Wind",
				capacity: 0,
				owned: 0,
				operating: 0,
				damage: 0,
				revenue: 0,
				cost: 0,
				icon: '<circle cx="500" cy="500" r="475" fill="#fff"/><path d="M497.7 439.1c-4.2-.2-8.3-1.1-12.3-2.6-1.2-.5-2.4-1.4-3.6-1.9L464.5 941h71.4l-16.1-470.8c-8.1-10.8-15.5-21-22.1-31.1z"/><path d="M595.6 467.9c-21.1-25.3-42.1-47.6-60.5-64.7 1.8-9 0-18.1-4.6-25.6 11.6-22.5 23.9-50.7 35.4-81.9 29.1-79.3 41.4-147.8 27.5-152.9-13.9-5.1-48.8 55.1-77.9 134.4-11.4 31.1-20.2 60.4-25.9 85-8.6 2.7-16.2 8.7-20.7 17-25.2 1.2-55.7 4.6-88.3 10.3-83.3 14.5-148.7 38-146.1 52.7 2.5 14.6 72.1 14.7 155.3.3 32.5-5.6 62.2-12.7 86.3-20 3.3 3.1 7.3 5.5 11.8 7.2 4.6 1.7 9.2 2.3 13.8 2.1 13.6 21.1 31.6 45.5 52.6 70.6 54.2 64.9 107.3 109.7 118.7 100.2 11.4-9.5-23.3-69.8-77.4-134.7zm-77.1-65.1c-3.7 10.1-14.9 15.3-25 11.6-10.1-3.7-15.3-14.9-11.6-25 3.7-10.1 14.9-15.3 25-11.6 10.1 3.8 15.3 15 11.6 25z"/><path d="M500 0C223.9 0 0 223.9 0 500s223.9 500 500 500 500-223.9 500-500S776.1 0 500 0zm0 940C257 940 60 743 60 500S257 60 500 60s440 197 440 440-197 440-440 440z"/>'
			},
			{
				name: "Hydroelectric",
				capacity: 0,
				owned: 0,
				operating: 0,
				damage: 0,
				revenue: 0,
				cost: 0,
				icon: '<circle cx="500" cy="500" r="475" fill="#fff"/><path d="M500 0C223.9 0 0 223.9 0 500s223.9 500 500 500 500-223.9 500-500S776.1 0 500 0zm303.1 181c-21.3-1.4-53-2.7-71.9 1.1-25.5 5.2-51.1-.7-75.4 1.4-118.5 9.9-78.8-49.4-149.3-38.1-70.5 11.3-74.7 5.6-117-12.7-42.3-18.3-46.5-2.8-69.1 7.1-5.1 2.2-9.6 1.4-17.3-3-4.7-2.7-16.2-11.1-24.4-17.2C343.7 81.7 419.3 60 500 60c117.5 0 224.2 46 303.1 121zm-599 644.7L323.7 552s-6.3-27.4 12.7-42.2l122.4 4.2c-17.9 12.6-8.5 46.4-8.5 46.4l-73.6 199.9 86.8-197.8s-10.9-29.4 13.2-47.9l123.3 4.2c-16.9 19-8.4 43.2-8.4 43.2L512 771.4 604.8 562s-11-30.9 11.6-42.6l111.3 4.1s-16.3 15.3-9.5 43.2L578.7 933c-25.5 4.6-51.9 7-78.7 7-113.9 0-217.7-43.3-295.9-114.3zm414.5 98.1L749 580.5V476.4L302.7 467l-.1-34.3-34.6-1.6-17.2 5.8h.1L296 439v93.3L178.2 800.1C109.7 726.7 66.1 629.9 60.6 523c62.6 1.8 125.8 3.9 186.4 6.3v-64l-185.3-4.1c2.9-33.6 9.6-66.1 19.7-97.1 51.2 2.9 125.1 5.9 191.1 4.3 115.6-2.9 277.8 16.9 153.7-28.3 0 0-49.4-62-258.1-95.9-3.3-.5-14.5-1.1-25-1.6 3.6-5 7.4-10 11.2-14.8 7.1 1.1 14.6 2.3 16 2.6 187 30.3 250.2 83.8 263.9 97.7 5 1.9 9.7 3.6 13.9 5.3 28.5 4.9 66.8 9.6 114.1 11.3 87.7 3.3 227.5.5 347.6 6.4 1.4.1 2.8.1 4.2.1 14.6 40.6 23.4 83.9 25.4 129l-139.7-3.1-.7-25.6-29.7-1.9-18.5 5.3h-1l40.1 3.2v94.8c53.5 2.3 104.5 5.5 145.7 8.5-24.2 174.4-150.8 316-317 362.4z"/>'
			}
		]
	},
	computed: {
		production: function() {
			var total = 0;
			clout.data.resources.forEach(function(element) {
				total += element.operating;
			});
			return total;
		},
		income: function() {
			var total = 0;
			clout.data.resources.forEach(function(element) {
				total += element.operating * element.revenue;
			});
			return total;
		},
		yearlyDamage: function() {
			var total = 0;
			clout.data.resources.forEach(function(element) {
				total += element.operating * element.damage;
			});
			return total;
		},
		energy: function() {
			var demand = clout.data.country.demand;
			var production = clout.computed.production();
			
			var base = production / demand * 100;
			var direction = "deficit";
			
			if (production > demand) {
				base = demand / production * 100;
				direction = "excess";
			}
			
			return { base: base + "%", extra: 100 - base + "%", direction: direction, difference: production - demand };
		}
	},
	methods: {
		login: function() {
			clout.data.connection.status = "connecting";
			
			if (clout.data.connection.creatingRoom) {
				request("create", "POST", { "username": clout.data.connection.name }, "", function(data, error) {
					if (error) {
						alert("Error: " + JSON.stringify(data));
					} else {
						console.log(data);
						clout.data.connection.invitePrompt = true;
						clout.data.connection.status = "connected";
						
						clout.data.connection.room = data.gameid;
						clout.data.connection.token = data.token;
						clout.data.country.name = data.country.name;
						
						clout.data.resources[0].capacity = data.country.coal.capacity;
						clout.data.resources[1].capacity = data.country.oil.capacity;
						clout.data.resources[2].capacity = data.country.gas.capacity;
						clout.data.resources[3].capacity = data.country.nuclear.capacity;
						clout.data.resources[4].capacity = data.country.geothermal.capacity;
						clout.data.resources[5].capacity = data.country.solar.capacity;
						clout.data.resources[6].capacity = data.country.wind.capacity;
						clout.data.resources[7].capacity = data.country.hydroelectric.capacity;
						
						clout.methods.update();
					}
				});
			} else {
				request("join", "POST", { "username": clout.data.connection.name, "gameid": clout.data.connection.room }, "", function(data, error) {
					if (error) {
						alert("Error: " + JSON.stringify(data));
					} else {
						clout.data.connection.status = "connected";
						clout.data.connection.token = data.token;
						clout.data.country.name = data.country.name;
						clout.methods.update();
					}
				});
			}
		},
		start: function() {
			clout.data.connection.invitePrompt = false;
			
			request("start", "POST", {}, clout.data.connection.token, function(data, error) {
				if (error) {
					alert("Error: " + JSON.stringify(data));
				} else {
					// Do anything when game starts?
				}
			});
		},
		update: function() {
			request("player", "GET", {}, clout.data.connection.token, function(data, error) {
				if (error) {
					alert("Error: " + JSON.stringify(data));
				} else {
					console.log(data);
					
					clout.data.resources[0].cost = data.player.coal.cost;
					clout.data.resources[1].cost = data.player.oil.cost;
					clout.data.resources[2].cost = data.player.gas.cost;
					clout.data.resources[3].cost = data.player.nuclear.cost;
					clout.data.resources[4].cost = data.player.geothermal.cost;
					clout.data.resources[5].cost = data.player.solar.cost;
					clout.data.resources[6].cost = data.player.wind.cost;
					clout.data.resources[7].cost = data.player.hydroelectric.cost;
					
					clout.data.country.damage = data.player.damage;
					clout.data.country.demand = data.player.demand / 1000000000;
					clout.data.country.money = data.player.money;
				}
				setTimeout(clout.methods.update, 1000);
			});
		},
		buy: function(index) {
			if (clout.data.country.money >= clout.data.resources[index].cost && (clout.data.resources[index].owned < clout.data.resources[index].capacity || clout.data.resources[index].capacity < 0)) {
				request("purchase/" + index, "POST", {}, clout.data.connection.token, function(data, error) {
					if (error) {
						alert("Error: " + JSON.stringify(data));
					} else {
						console.log(data);
						
						clout.data.resources[0].operating = data.player.coal.operational;
						clout.data.resources[0].owned = data.player.coal.owned;
						
						clout.data.resources[1].operating = data.player.oil.operational;
						clout.data.resources[1].owned = data.player.oil.owned;
						
						clout.data.resources[2].operating = data.player.gas.operational;
						clout.data.resources[2].owned = data.player.gas.owned;
						
						clout.data.resources[3].operating = data.player.nuclear.operational;
						clout.data.resources[3].owned = data.player.nuclear.owned;
						
						clout.data.resources[4].operating = data.player.geothermal.operational;
						clout.data.resources[4].owned = data.player.geothermal.owned;
						
						clout.data.resources[5].operating = data.player.solar.operational;
						clout.data.resources[5].owned = data.player.solar.owned;
						
						clout.data.resources[6].operating = data.player.wind.operational;
						clout.data.resources[6].owned = data.player.wind.owned;
						
						clout.data.resources[7].operating = data.player.hydroelectric.operational;
						clout.data.resources[7].owned = data.player.hydroelectric.owned;
						
						clout.data.country.money = data.player.money;
					}
				});
			}
		},
		setOperational: function(index) {
			request("operational/" + index, "POST", { "number": parseInt(document.querySelectorAll(".resources .resource")[index].querySelector(".owned input").value) }, clout.data.connection.token, function(data, error) {
				if (error) {
					alert("Error: " + JSON.stringify(data));
				} else {
					// Do nothing?
				}
			});
		}
	},
	filters: {
		formatMoney: function(value) {
			var base = Math.floor(Math.log(Math.abs(value)) / Math.log(1000));
			var suffix = "kmbtqQsSondUDT"[base - 1];
			if (suffix) {
				return "$" + Math.round(value / Math.pow(1000, base) * 100) / 100 + suffix;
			} else {
				return "$" + value;
			}
		}
	}
};

document.addEventListener("DOMContentLoaded", function() {
	vue = new Vue(clout);
});

var request = function(path, method, body, token, callback) {
	var root = "https://cloutgame.herokuapp.com/api";
	if (path[0] !== "/") {
		path = "/" + path;
	}
	
	var request = new XMLHttpRequest();
	request.open(method.toUpperCase(), root + path, true);
	request.setRequestHeader("Content-Type", "application/json");
	request.setRequestHeader("Authorization", token);
	
	request.onload = function() {
		if (request.status >= 200 && request.status < 400) {
			callback(JSON.parse(request.responseText), false);
		} else {
			try {
				callback(JSON.parse(request.responseText), true);
			} catch(e) {
				callback({}, true);
			}
		}
	};
	
	request.onerror = function() {
		callback({}, true);
	};
	
	if (method.toUpperCase() === "GET") {
		request.send();
	} else if (method.toUpperCase() === "POST") {
		request.send(JSON.stringify(body));
	}
};
