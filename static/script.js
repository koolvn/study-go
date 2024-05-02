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
    }).then(response => {
        const statusMessage = getStatusMessage(response.status);
        if (response.headers.get('Content-Type').includes('application/json')) {
            return response.json().then(json_response => {
                // Check if details is an object and stringify if true
                let details = json_response.details;
                if (typeof details === 'object' && details !== null) {
                    // Optionally, add indentation for better readability
                    details = JSON.stringify(details, null, 2);
                }
                return {
                    status: statusMessage,
                    details: details || 'No details provided', // Fallback message
                };
            });
        } else {
            return {
                status: statusMessage,
                details: 'Response was not in JSON format',
            };
        }
    })
        .then(response => {
            console.log('JSON RESPONSE: ', response);
            document.getElementById('status').textContent = response.status;
            // Use innerText if you've formatted the JSON for readability
            document.getElementById('details').innerText = response.details;
        })
        .catch(e => {
            console.log('There was a problem with the fetch operation: ' + e.message);
            document.getElementById('status').textContent = 'Error performing the request';
            document.getElementById('details').textContent = e.message;
        });
});

function getStatusMessage(statusCode) {
    const statusMessages = {
        200: '200 OK',
        400: '400 Bad Request',
        401: '401 Unauthorized',
        403: '403 Forbidden',
        404: '404 Not Found',
        500: '500 Internal Server Error',
        // Add more mappings as needed
    };

    return statusMessages[statusCode] || `${statusCode} Unknown Status`;
}
