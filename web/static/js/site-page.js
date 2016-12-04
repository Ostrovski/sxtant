(function(ns) {

	ns.Sextant = ns.Sextant || {};

	var questions = [];  // sorted by creation_date asc
	var queue = [];
	var tagsInclude = ['python', 'flask', 'go', 'php', 'javascript'];
	var tagsExclude = ['css', 'html', 'html5', 'angularjs', 'jquery'];

	var questionsList = document.getElementsByClassName('js-questions')[0];

	Sextant.run = function() {
		Notification.requestPermission(function(permission) {
        	if (permission !== "granted") {
            	alert("Notifications disabled");
            	return;
        	}

        	setInterval(function() {
				if (queue.length > 3 && _now() - queue[0].createdAt >= 5) {
					queue[0].notification.close();
					queue.shift();
				}
        	}, 1000);

			var ws = new WebSocket('ws://' + window.location.host + '/sock');
			ws.onmessage = function(event) {
				var newQuestions = JSON.parse(event.data).items;
				for (var i = newQuestions.length - 1; i >= 0 ; i--) {
					var q = newQuestions[i];
					if (0 === questions.length || q.creation_date > questions[questions.length - 1].creation_date) {
						if (_isInteresting(q)) {
							var qId = 'q' + q.question_id;
							var notification = _questionNotification(q);
							queue.push({notification: notification, createdAt: _now()});
							questions.push(q);

							notification.onclick = function() {
								var prev = document.getElementsByClassName('current-q');
								for (var i = 0; i < prev.length; i++) {
									prev[i].setAttribute('class', '');
								}

								window.focus();
								document.getElementById(qId).setAttribute('class', 'current-q');
								document.getElementById('link-' + qId).click();
							};

							var questionEl = document.createElement('li');
							questionEl.setAttribute('id', qId);

							var linkEl = document.createElement('a');
							linkEl.setAttribute('id', 'link-' + qId);
							linkEl.setAttribute('href', q.link);
							linkEl.setAttribute('target', '_blank');
							linkEl.innerHTML = q.title;

							var tagsEl = document.createElement('span');
							tagsEl.setAttribute('class', 'q-tags');
							tagsEl.innerHTML = q.tags.join(' ');

							questionEl.appendChild(linkEl);
							questionEl.appendChild(tagsEl);
							questionsList.appendChild(questionEl);
						}
					}
				}
			};
    	});
	};

	function _questionNotification(q) {
		return new Notification(q.tags.join(' '), {body: q.title});
	}

	function _now() {
		return parseInt(+new Date()/1000);
	}

	function _isInteresting(q) {
		for (var i = 0; i < tagsExclude.length; i++) {
			if (-1 !== q.tags.indexOf(tagsExclude[i])) {
				return false;
			}
		}

		for (var j = 0; j < tagsInclude.length; j++) {
			if (-1 !== q.tags.indexOf(tagsInclude[j])) {
				return true;
			}
		}

		return false;
	}
})(window);
