document.addEventListener("DOMContentLoaded", function () {
    // ====== INDEX - ENVIAR DINERO ======
    const enviarBtn = document.querySelector("#enviarModal .btn-primary");
    if (enviarBtn) {
        enviarBtn.addEventListener("click", function () {
            const cuentaDestino = document.getElementById("cuentaDestino");
            const monto = document.getElementById("montoEnviar");

            if (!cuentaDestino || !monto || !cuentaDestino.value || !monto.value) {
                alert("Por favor completa todos los campos.");
                return;
            }

            alert(`Enviando ${monto.value} a la cuenta ${cuentaDestino.value}`);
            // Aquí puedes agregar la lógica para hacer la petición al backend en Go

            // Cerrar modal automáticamente
            $('#enviarModal').modal('hide');
        });
    }
    
    // ====== PERSONAS ======
    
    // Manejar apertura del modal de edición para personas
    $('#editarPersonaModal').on('show.bs.modal', function (event) {
        var button = $(event.relatedTarget);
        var id = button.data('id');
        var nombre1 = button.data('nombre1');
        var nombre2 = button.data('nombre2');
        var apellido1 = button.data('apellido1');
        var apellido2 = button.data('apellido2');
        var documento = button.data('documento');
        var correo = button.data('correo');
        var direccion = button.data('direccion');
        var telefono = button.data('telefono');

        var modal = $(this);
        modal.find('#editId').val(id);
        modal.find('#editPrimerNombre').val(nombre1);
        modal.find('#editSegundoNombre').val(nombre2);
        modal.find('#editPrimerApellido').val(apellido1);
        modal.find('#editSegundoApellido').val(apellido2);
        modal.find('#editDocumento').val(documento);
        modal.find('#editCorreo').val(correo);
        modal.find('#editDireccion').val(direccion);
        modal.find('#editTelefono').val(telefono);
    });

    // Función eliminar persona
    window.eliminarPersona = function(id) {
        if (confirm('¿Está seguro de eliminar esta persona?')) {
            window.location.href = '/personas/eliminar?id=' + id;
        }
    };
    
    // Manejar apertura del modal de edición para clientes
    $('#editarClienteModal').on('show.bs.modal', function (event) {
        var button = $(event.relatedTarget);
        var id = button.data('id');
        var idPersona = button.data('idpersona');
        var tipo = button.data('tipo');
        var saldo = button.data('saldo');

        var modal = $(this);
        modal.find('#editId').val(id);
        modal.find('#editIdPersona').val(idPersona);
        modal.find('#editTipoCuenta').val(tipo);
        modal.find('#editSaldo').val(saldo);
    });

    // Función eliminar cliente
    window.eliminarCliente = function(id) {
        if (confirm('¿Está seguro de eliminar este cliente?')) {
            window.location.href = '/clientes/eliminar?id=' + id;
        }
    };

    // Manejar apertura del modal de VER detalles para transacciones (solo lectura)
    $('#verTransaccionModal').on('show.bs.modal', function (event) {
        var button = $(event.relatedTarget);
        var id = button.data('id');
        var idCliente = button.data('idcliente');
        var cliente = button.data('cliente');
        var tipo = button.data('tipo');
        var monto = button.data('monto');
        var fecha = button.data('fecha');
        var descripcion = button.data('descripcion');

        var modal = $(this);
        modal.find('#verIdTransaccion').text(id);
        modal.find('#verIdCliente').text(idCliente);
        modal.find('#verNombreCliente').text(cliente || 'Cliente #' + idCliente);
        modal.find('#verTipo').html('<span class="badge badge-' + getTipoBadge(tipo) + '">' + tipo + '</span>');
        modal.find('#verMonto').text('$' + parseFloat(monto).toFixed(2));
        modal.find('#verFecha').text(fecha);
        modal.find('#verDescripcion').text(descripcion || 'Sin descripción');
    });

    // Función auxiliar para obtener el color del badge según el tipo
    function getTipoBadge(tipo) {
        if (tipo === 'Depósito') return 'success';
        if (tipo === 'Retiro') return 'warning';
        return 'info';
    }

    // Función eliminar transacción
    window.eliminarTransaccion = function(id) {
        if (confirm('¿Está seguro de eliminar esta transacción? Esta acción no se puede deshacer.')) {
            window.location.href = '/transacciones/eliminar?id=' + id;
        }
    };

    // Cargar clientes en el select del modal de crear transacción
    function cargarClientes() {
        $.ajax({
            url: '/api/clientes',
            method: 'GET',
            dataType: 'json',
            success: function(data) {
                var select = $('#selectCliente');
                var selectDestino = $('#selectClienteDestino');
                
                select.empty();
                selectDestino.empty();
                
                if (data.length === 0) {
                    select.append('<option value="">No hay clientes disponibles</option>');
                    selectDestino.append('<option value="">No hay clientes disponibles</option>');
                } else {
                    select.append('<option value="">Seleccione un cliente</option>');
                    selectDestino.append('<option value="">Seleccione el cliente destino</option>');
                    data.forEach(function(cliente) {
                        select.append('<option value="' + cliente.id + '">' + cliente.nombre + '</option>');
                        selectDestino.append('<option value="' + cliente.id + '">' + cliente.nombre + '</option>');
                    });
                }
            },
            error: function() {
                $('#selectCliente').html('<option value="">Error al cargar clientes</option>');
                $('#selectClienteDestino').html('<option value="">Error al cargar clientes</option>');
            }
        });
    }

    // Mostrar/ocultar campos de transferencia según el tipo seleccionado
    $(document).on('change', '#tipoTransaccion', function() {
        var tipo = $(this).val();
        console.log('Tipo seleccionado:', tipo); // Para debugging
        
        if (tipo === 'Transferencia') {
            $('#camposTransferencia').slideDown();
            $('#selectClienteDestino').prop('required', true);
        } else {
            $('#camposTransferencia').slideUp();
            $('#selectClienteDestino').prop('required', false);
            $('#selectClienteDestino').val('');
        }
    });

    // Cargar clientes cuando se abre el modal de crear transacción
    $('#crearTransaccionModal').on('show.bs.modal', function() {
        cargarClientes();
        // Resetear el formulario
        $('#tipoTransaccion').val('');
        $('#camposTransferencia').hide();
        $('#selectClienteDestino').prop('required', false);
    });
});
