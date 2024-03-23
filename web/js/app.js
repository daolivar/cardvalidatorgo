// document.getElementById("cardForm") is returning null
document.getElementById("cardForm").addEventListener("submit", function (event) {
    event.preventDefault(); // Prevent form submission

    var cardNumber = document.getElementById("cardnumber").value; // Get the card number from the input field

    // Create a new XMLHttpRequest object
    var xhr = new XMLHttpRequest();

    // Configure the request
    xhr.open("POST", "/validate", true);
    xhr.setRequestHeader("Content-Type", "application/json");

    // Set up the onload callback function
    xhr.onload = function () {
        if (xhr.status === 200) {
            // If the request is successful, display the response
            document.getElementById("validationResult").innerHTML = xhr.responseText;
        } else {
            // If there's an error, display an error message
            document.getElementById("validationResult").innerHTML = "Error: " + xhr.statusText;
        }
    };

    // Set up the onerror callback function
    xhr.onerror = function () {
        // If there's a network error, display an error message
        document.getElementById("validationResult").innerHTML = "Network Error";
    };

    // Convert the card number to JSON format and send the request
    xhr.send(JSON.stringify({ "cardnumber": cardNumber }));
});
