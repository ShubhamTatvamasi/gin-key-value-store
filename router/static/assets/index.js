// Set Key Value Pair
async function setKeyValue() {

    let formData = new FormData();

    let setNewKey = document.getElementById('setNewKey').value
    let setNewValue = document.getElementById('setNewValue').value
    
    formData.append('key', setNewKey);
    formData.append('value', setNewValue);

    await fetch('/post', { method: 'POST', body: formData })
}

// Get value from key
async function getValue() {

    let keyFromUser = document.getElementById('newKey').value

    let getNewKey = await fetch('/get?key=' + keyFromUser);
    let newKey = await getNewKey.text();

    document.getElementById("key").innerHTML = newKey;
}

// Subscribe New Key
async function subscribeKey() {

    let formData = new FormData();

    let subscribeNewKey = document.getElementById('subscribeKey').value

    formData.append('key', subscribeNewKey);

    await fetch('/subscribe', { method: 'POST', body: formData })
}

// Unsubscribe a Key
async function unsubscribeKey() {

    let formData = new FormData();

    let unsubscribeKey = document.getElementById('unsubscribeKey').value

    formData.append('key', unsubscribeKey);

    await fetch('/unsubscribe', { method: 'POST', body: formData })
}

// Create a new EventSource for a Live Subscriptions
let source = new EventSource('/stream');

// Function for updating subscriptions
source.onmessage = (e) => {

    // parsing json string to json 
    let subscriptions = JSON.parse(e.data);

    // Creating new table for subscribed key value pairs
    let table = "<table>";

    for (let k in subscriptions ) {
        table += "<tr><td>"+k+"</td><td>"+subscriptions[k]+"</td></tr>";
    }
    table += "</table>"

    // Updating the subscriptions table
    document.getElementById("subscriptions").innerHTML = table;

}
