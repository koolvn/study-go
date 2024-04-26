document.getElementById('auth-form').addEventListener('submit', function (event) {
    event.preventDefault();

    let username = document.getElementById('username').value;
    let password = document.getElementById('password').value;

    fetch('/auth', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({username: username, password: password}),
    }).then(response => response.json())
        .then(json_response => {
            console.log('JSON RESPONSE: ', json_response)
            let error_message = ''
            if (json_response['status'] === 'Unauthorized') {
                error_message = json_response['error']
            }
            document.getElementById('status').textContent = json_response['status'];
            document.getElementById('error-message').textContent = error_message;
        })
        .catch(e => {
            console.log('There was a problem with the fetch operation: ' + e.message);
        });
});
