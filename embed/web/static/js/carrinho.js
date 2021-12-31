
$(document).ready(function () {


	showLoaderGeneric(function () {

		setTimeout(function () {

			// Ao carregar a pagina, verificar se tem 
			// dados do usuario na memoria

			var nome = LocalStorage.get('gcloud.poc.user_nome', nome);
			var cep = LocalStorage.get('gcloud.poc.endereco_cep', cep);
			var uf = LocalStorage.get('gcloud.poc.endereco_uf', uf);
			var cidade = LocalStorage.get('gcloud.poc.endereco_cidade', cidade);
			var bairro = LocalStorage.get('gcloud.poc.endereco_bairro', bairro);
			var logradouro = LocalStorage.get('gcloud.poc.endereco_logradouro', logradouro);
			var numero = LocalStorage.get('gcloud.poc.endereco_numero', numero);
			var complemento = LocalStorage.get('gcloud.poc.endereco_complemento', complemento);

			var checkUser = 	( nome !== null && nome.length > 0 ) 
							&&	( cep !== null && cep.length > 0 )
							&&	( uf !== null && uf.length > 0 )
							&&	( cidade !== null && cidade.length > 0 )
							&&	( bairro !== null && bairro.length > 0 )
							&&	( logradouro !== null && logradouro.length > 0 )
							&&	( numero !== null && numero.length > 0 && numero > 0 )
						;


			// Caso ja tenha dados do usuario, 
			// redirecionar para o finalizar compra

			if (checkUser) {

		    	window.location.href = 'finalizar-compra.html';

			}

			hideLoaderGeneric();

		}, 100);
	});

	// Mascara para os campos de CEP
	$('.form-control.form-control-cep').formatter({
		'pattern': '{{99}}.{{999}}-{{999}}',
		'persistent': false
	});

	// Mascara para os campos de Numero do endereco
	$('.form-control.form-control-numero').formatter({
		'pattern': '{{9999999999}}',
		'persistent': false
	});

});


// Buscar CEP
$(document).on('click', '.btn-buscar-cep', function () {
	
	var nome = $('#txtNome').val().trim();
	var cep = $('#txtCep').val().trim();

	var isValid = true;

	if (!nome || nome.length == 0) {
		
		$('#txtNome').addClass('form-control-danger');
		$('#txtNome').parent('.form-group').addClass('has-danger');
		
		isValid = false;

	} else {
		$('#txtNome').removeClass('form-control-danger');
		$('#txtNome').parent('.form-group').removeClass('has-danger');		
	}

	if (!cep || cep.length == 0) {
		
		$('#txtCep').addClass('form-control-danger');
		$('#txtCep').parent('.form-group').addClass('has-danger');
		
		isValid = false;

	} else {
		$('#txtCep').removeClass('form-control-danger');
		$('#txtCep').parent('.form-group').removeClass('has-danger');		
	}

	if (!isValid) {

		toastr.error('Por favor, preencha os campos indicados!');

	} else {

		cep = cep.replace('.', '');
		cep = cep.replace('-', '');

		var path = 'https://raw.githubusercontent.com/jeffotoni/api.cep/master/v1/cep/' + cep;

		showLoaderButton($('.btn-buscar-cep'), function () {

			AxiosJson('GET', path, null, {}, function (data, status) {

				if (status != 200) {

					toastr.error('O CEP n&atilde;o foi encontrado ou n&atilde;o &eacute; v&aacute;lido!');

				} else {

					var obj = data;

					if (obj.uf !== null && obj.uf !== undefined) {
						$('#txtCepUf').val(obj.uf);
					}

					if (obj.cidade !== null && obj.cidade !== undefined) {
						$('#txtCepCidade').val(obj.cidade);
					}

					if (obj.bairro !== null && obj.bairro !== undefined) {
						$('#txtCepBairro').val(obj.bairro);
					}

					if (obj.logradouro !== null && obj.logradouro !== undefined) {
						$('#txtCepLogradouro').val(obj.logradouro);
					}

					$('#FormCep').css('visibility', 'visible');

				}

				hideLoaderButton($('.btn-buscar-cep'));

			});

		});

	}
});


// Carregar página de finalizar compra
$(document).on('click', '.btn-finalizar-compra', function () {

	var nome = $('#txtNome').val().trim();
	var cep = $('#txtCep').val().trim();
	var uf = $('#txtCepUf').val().trim();
	var cidade = $('#txtCepCidade').val().trim();
	var bairro = $('#txtCepBairro').val().trim();
	var logradouro = $('#txtCepLogradouro').val().trim();
	var numero = $('#txtCepNumero').val().trim();
	var complemento = $('#txtCepComplemento').val().trim();

	var isValid = true;

	if (!nome || nome.length == 0) {
		
		$('#txtNome').addClass('form-control-danger');
		$('#txtNome').parent('.form-group').addClass('has-danger');
		
		isValid = false;

	} else {
		$('#txtNome').removeClass('form-control-danger');
		$('#txtNome').parent('.form-group').removeClass('has-danger');		
	}

	if (!cep || cep.length == 0) {
		
		$('#txtCep').addClass('form-control-danger');
		$('#txtCep').parent('.form-group').addClass('has-danger');
		
		isValid = false;

	} else {
		$('#txtCep').removeClass('form-control-danger');
		$('#txtCep').parent('.form-group').removeClass('has-danger');		
	}

	if (!uf || uf.length == 0) {
		
		$('#txtCepUf').addClass('form-control-danger');
		$('#txtCepUf').parent('.form-group').addClass('has-danger');
		
		isValid = false;

	} else {
		$('#txtCepUf').removeClass('form-control-danger');
		$('#txtCepUf').parent('.form-group').removeClass('has-danger');		
	}

	if (!cidade || cidade.length == 0) {
		
		$('#txtCepCidade').addClass('form-control-danger');
		$('#txtCepCidade').parent('.form-group').addClass('has-danger');
		
		isValid = false;

	} else {
		$('#txtCepCidade').removeClass('form-control-danger');
		$('#txtCepCidade').parent('.form-group').removeClass('has-danger');		
	}

	if (!bairro || bairro.length == 0) {
		
		$('#txtCepBairro').addClass('form-control-danger');
		$('#txtCepBairro').parent('.form-group').addClass('has-danger');
		
		isValid = false;

	} else {
		$('#txtCepBairro').removeClass('form-control-danger');
		$('#txtCepBairro').parent('.form-group').removeClass('has-danger');		
	}

	if (!logradouro || logradouro.length == 0) {
		
		$('#txtCepLogradouro').addClass('form-control-danger');
		$('#txtCepLogradouro').parent('.form-group').addClass('has-danger');
		
		isValid = false;

	} else {
		$('#txtCepLogradouro').removeClass('form-control-danger');
		$('#txtCepLogradouro').parent('.form-group').removeClass('has-danger');		
	}

	if (!numero || numero.length == 0) {
		
		$('#txtCepNumero').addClass('form-control-danger');
		$('#txtCepNumero').parent('.form-group').addClass('has-danger');
		
		isValid = false;

	} else {
		$('#txtCepNumero').removeClass('form-control-danger');
		$('#txtCepNumero').parent('.form-group').removeClass('has-danger');		
	}


	if (!isValid) {

		toastr.error('Por favor, preencha os campos indicados!');

	} else {

		LocalStorage.set('gcloud.poc.user_nome', nome);
		LocalStorage.set('gcloud.poc.endereco_cep', cep);
		LocalStorage.set('gcloud.poc.endereco_uf', uf);
		LocalStorage.set('gcloud.poc.endereco_cidade', cidade);
		LocalStorage.set('gcloud.poc.endereco_bairro', bairro);
		LocalStorage.set('gcloud.poc.endereco_logradouro', logradouro);
		LocalStorage.set('gcloud.poc.endereco_numero', numero);
		LocalStorage.set('gcloud.poc.endereco_complemento', complemento);

    	window.location.href = 'finalizar-compra.html';

	}
});