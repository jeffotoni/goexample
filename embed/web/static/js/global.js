
$(document).ready(function () {

	// Configuração das mensagens de alerta no topo
	toastr.options = {
		"closeButton": false,
		"debug": false,
		"newestOnTop": false,
		"progressBar": false,
		"positionClass": "toast-top-full-width",
		"preventDuplicates": false,
		"onclick": null,
		"showDuration": "300",
		"hideDuration": "1000",
		"timeOut": "2500",
		"extendedTimeOut": "1000",
		"showEasing": "swing",
		"hideEasing": "linear",
		"showMethod": "fadeIn",
		"hideMethod": "fadeOut"
	}


	// Validação de vazio em campos obrigatórios
	$(document).on('blur', 'input.form-required, textarea.form-required, select.form-required', function () {
		var el = $(this);
		var valor = el.val();

		if (valor !== null && valor !== undefined) {

			if (valor.trim().length == 0) {

				el.addClass('form-control-danger');
				el.parent('.form-group').addClass('has-danger');
				el.parent('.input-group').parent('.form-group').addClass('has-danger');
				el.parent('div').parent('.form-group').addClass('has-danger');
				el.parent('div').parent('div').parent('.form-group').addClass('has-danger');

			} else {

				el.removeClass('form-control-danger');
				el.parent('.form-group').removeClass('has-danger');
				el.parent('.input-group').parent('.form-group').addClass('has-danger');
				el.parent('div').parent('.form-group').removeClass('has-danger');
				el.parent('div').parent('div').parent('.form-group').removeClass('has-danger');

			}
		}
	});

});




// Funcao para as chamadas axios
const AxiosJson = function(metodoEnvio, path, params, objHeaders, callback) {

    if (path) {
        
        axios({
            timeout: 10000,
            url: path,
            headers: objHeaders, 
            data: JSON.stringify(params),
            responseType: "json",
            method: metodoEnvio
        })
        .then(function (response) {

            if (response.status !== null && response.status !== undefined && response.status === 200) {

                callback(response.data, response.status);

            } else {

                toastr.error("Erro: o servidor n&atilde;o respondeu corretamente.");

            }

        })
        .catch(function (error, response) {

            if (error.code !== null && error.code !== undefined && error.code === 'ECONNABORTED' && callback) {

                // E CONN ABORTED -> Timeout
                // Tratar o timeout em cada chamada 
                // Esconder loader, destravar botão...
                
                callback(error.code, error.response.status);

            } else if (error.response !== null && error.response !== undefined && error.response) {

                // The request was made and the server responded with a status code
                // that falls out of the range of 2xx

                // console.log(error.response.data);
                // console.log(error.response.status);
                // console.log(error.response.headers);

            } else if (error.request !== null && error.request !== undefined && error.request) {

                // The request was made but no response was received
                // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
                // http.ClientRequest in node.js

                console.log("Request error: ", error.request);

            } else {

                // Something happened in setting up the request that triggered an Error
                
                if (error.message !== null && error.message !== undefined) {
                	
                	console.log('Error', error.message);

                } else {

                	console.log('Error na request!');

                }

            }

            if (callback && error.response !== null && error.response !== undefined) {
            
                callback(error.response.data, error.response.status);
            
            }
        
        });

    } else {

        toastr.error("Erro: o caminho n&atilde;o foi encontrado.");

    }
};


var hideLoaderButton = function(elButton) {
    $(elButton).find(".loader").fadeOut(100, function () {
        $(elButton).removeAttr("disabled").find(".text").fadeIn(100);
    });
};

var hideLoaderGeneric = function() {
	$('#shadowWithLoader').removeClass('show').hide();
}


var showLoaderButton = function(elButton, callback) {
    $(elButton).blur().find(".text").fadeOut(100, function () {
        $(elButton).attr("disabled", "disabled").find(".loader").fadeIn(100, function () {
            setTimeout(function () {
                if (callback)
                callback();
            }, 1000);
        });
    });
};

var showLoaderGeneric = function(callback) {
    $('#shadowWithLoader').addClass('show').fadeIn(1, function () {
	    if (callback) {
	        callback();
	    }
    });
};



// Limpa os dados do usuario do LocalStorage 
// e retorna para a pagina de carrinho

$(document).on('click', '.btn-sair', function () {

	LocalStorage.remove('gcloud.poc.user_nome');
	LocalStorage.remove('gcloud.poc.endereco_cep');
	LocalStorage.remove('gcloud.poc.endereco_uf');
	LocalStorage.remove('gcloud.poc.endereco_cidade');
	LocalStorage.remove('gcloud.poc.endereco_bairro');
	LocalStorage.remove('gcloud.poc.endereco_logradouro');
	LocalStorage.remove('gcloud.poc.endereco_numero');
	LocalStorage.remove('gcloud.poc.endereco_complemento');

	window.location.href = 'index.html';
});