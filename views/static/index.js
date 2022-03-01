window.onload = function() {
    document.getElementById("input").focus()

    const messagesDiv = document.getElementById("messages")
    messagesDiv.scrollTop = messagesDiv.scrollHeight;
};
