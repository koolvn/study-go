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
            document.getElementById('status').textContent = json_response['status'];
            document.getElementById('details').textContent = json_response['details'];
        })
        .catch(e => {
            console.log('There was a problem with the fetch operation: ' + e.message);
        });
});
