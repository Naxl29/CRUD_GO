document.addEventListener("DOMContentLoaded", function () {
    const enviarBtn = document.querySelector(".btn-primary");
    enviarBtn.addEventListener("click", function () {
        const cuentaDestino = document.getElementById("cuentaDestino").value;
        const monto = document.getElementById("montoEnviar").value;

        if (!cuentaDestino || !monto) {
            alert("Por favor completa todos los campos.");
            return;
        }

        alert(`Enviando ${monto} a la cuenta ${cuentaDestino}`);
        // Aquí puedes agregar la lógica para hacer la petición al backend en Go

        // Cerrar modal automáticamente
        $('#enviarModal').modal('hide');
    });
});
