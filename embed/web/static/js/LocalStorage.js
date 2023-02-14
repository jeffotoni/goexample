
const LocalStorage = {

	set: (keyName, objValue) => {
		if (keyName && objValue) {

			if (!window.localStorage) {
				return null;
			}

			localStorage.setItem(keyName, JSON.stringify(objValue));
		}
	},

	get: (keyName) => {
		if (keyName) {

			if (!window.localStorage) {
				return null;
			}

			return JSON.parse(localStorage.getItem(keyName));
		}
		return null;
	},
	
	remove: (keyName) => {
		if (keyName) {

			if (!window.localStorage) {
				return null;
			}

			localStorage.removeItem(keyName);
		}
		return null;
	}

};