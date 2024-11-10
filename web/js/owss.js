// Event-Listener für das Formular, um das Standard-Submit-Verhalten zu verhindern
document.getElementById('owss_logon_formular').addEventListener('submit', function(event) {
    // Verstecke das Login-Formular
    document.getElementById("owss_logon_formular").style.display = "none";

    // Zeige das Element für die Login-Verarbeitung an
    document.getElementById("process_login").style.display = "block";

    // Verhindert das Neuladen der Seite
    event.preventDefault();

    // Erstelle FormData-Objekt und sende Daten mit fetch
    const formData = new FormData(event.target);

    fetch('/your-server-endpoint', {
        method: 'POST',
        body: formData
    })
    .then(response => {
        if (response.ok) {
            return response.json();
        } else {
            throw new Error('Network response was not ok');
        }
    })
    .then(data => {
        console.log('Success:', data);
        // Hier kannst du zusätzliche Logik für die erfolgreiche Antwort hinzufügen
    })
    .catch(error => {
        console.error('Error:', error);
    });
});