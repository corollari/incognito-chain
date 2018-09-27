function numberWithCommas(x) {
    return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}

window.onload = function () {
    var xhr = new XMLHttpRequest();   // new HttpRequest instance
    xhr.open("POST", api_url);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.onreadystatechange = function (oEvent) {
        if (xhr.status === 200) {
            console.log(this.responseText.toString());
            var response = JSON.parse(this.responseText.toString());
            if (response.Result != null) {
                document.getElementById("lb_publicKey").innerText = response.Result.PublicKey;
                document.getElementById("lb_readonlyKey").innerText = response.Result.ReadonlyKey;
                document.getElementById("loader").style.display = "none";
                document.getElementById("myDiv").style.display = "block";
                dumpprivkey(response.Result.PublicKey)
                getbalance();
            } else {
                if (response.Error != null) {
                    alert(response.Error.message);
                } else {
                    alert('Bad response');
                }
            }
        } else {
            alert('Network error');
        }
    };
    var url = new URL(window.location.href);
    var account = url.searchParams.get("account");
    xhr.send(JSON.stringify({
        jsonrpc: "1.0",
        method: "getaccountaddress",
        params: account,
        id: 1
    }));

    document.getElementById("bt_send").onclick = function () {
        sendmany()
    };

};

function dumpprivkey(publicKey) {
    var xhr = new XMLHttpRequest();   // new HttpRequest instance
    xhr.open("POST", api_url);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.onreadystatechange = function (oEvent) {
        if (xhr.status === 200) {
            console.log(this.responseText.toString());
            var response = JSON.parse(this.responseText.toString());
            if (response.Result != null) {
                document.getElementById("lb_privateKey").innerText = response.Result.PrivateKey;
            } else {
                if (response.Error != null) {
                    alert(response.Error.message);
                } else {
                    alert('Bad response');
                }
            }
        } else {
            alert('Network error');
        }
    };
    xhr.send(JSON.stringify({
        jsonrpc: "1.0",
        method: "dumpprivkey",
        params: publicKey,
        id: 1
    }));
}

function getbalance() {
    var url = new URL(window.location.href);
    var account = url.searchParams.get("account");
    var passphrase = window.localStorage['cash_passphrase'];

    var xhr = new XMLHttpRequest();   // new HttpRequest instance
    xhr.open("POST", api_url);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.onreadystatechange = function (oEvent) {
        if (xhr.status === 200) {
            console.log(this.responseText.toString());
            var response = JSON.parse(this.responseText.toString());
            if (response.Result != null) {
                document.getElementById("lb_balance").innerText = numberWithCommas(response.Result);
            } else {
                if (response.Error != null) {
                    alert(response.Error.message);
                } else {
                    alert('Bad response');
                }
            }
        } else {
            alert('Network error');
        }
    };
    xhr.send(JSON.stringify({
        jsonrpc: "1.0",
        method: "getbalance",
        params: [account, 1, passphrase],
        id: 1
    }));
}

function showLoading(show) {
    if (show) {
        document.getElementById("loader").style.display = "block";
        document.getElementById("myDiv").style.display = "none";
    } else {
        document.getElementById("loader").style.display = "none";
        document.getElementById("myDiv").style.display = "block";
    }
}

function sendmany() {
    var priKey = document.getElementById("lb_privateKey").innerText;
    var pubKey = document.getElementById("txt_address").value;
    var amount = document.getElementById("txt_amount").value;

    showLoading(true);

    var xhr = new XMLHttpRequest();   // new HttpRequest instance
    xhr.open("POST", api_url);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.onreadystatechange = function (oEvent) {
        showLoading(false);

        if (xhr.status === 200) {
            console.log(this.responseText.toString());
            var response = JSON.parse(this.responseText.toString());
            if (response.Result != null && response.Result != '') {
            } else {
                if (response.Error != null) {
                    alert(response.Error.message)
                } else {
                    alert('Bad response');
                }
            }
        } else {
            alert('Network error')
        }
    };
    var dest = {};
    dest[pubKey] = parseInt(amount);
    xhr.send(JSON.stringify({
        jsonrpc: "1.0",
        method: "sendmany",
        params: [priKey, dest, -1, 1],
        id: 1
    }));
}