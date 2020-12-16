
// Get value from key
async function getKey() {

    let keyFromUser = document.getElementById('newKey').value

    let getNewKey = await fetch('/get?key=' + keyFromUser);
    let newKey = await getNewKey.text();

    document.getElementById("key").innerHTML = newKey;
}

// Create a new EventSource for a live stream
let source = new EventSource('/stream/');

// Function for updating subscriptions
source.onmessage = (e) => {

    // parsing json string to json 
    let subscriptions = JSON.parse(e.data);

    // Creating new table for subscribed key value pairs
    let table = "<table>";

    for (var k in subscriptions ) {
        table += "<tr><td>"+k+"</td><td>"+subscriptions[k]+"</td></tr>";
    }
    table += "</table>"

    // Updating the subscriptions table
    document.getElementById("subscriptions").innerHTML = table;

}
