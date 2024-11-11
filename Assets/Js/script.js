document.querySelector("form").addEventListener("submit", function(event) {
    event.preventDefault();
    alert("Thank you for your message!");
    // You can further handle form submission via AJAX or another method if needed.
});
