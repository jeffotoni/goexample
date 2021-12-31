
$(document).ready(function () {

	showLoaderGeneric(function () {

		setTimeout(function () {

			// Ao carregar a pagina, verificar se tem 
			// dados do usuario na memoria

			var nome = LocalStorage.get('gcloud.poc.user_nome', nome);

			if (!nome) {

				$('.nav-item-sair').hide();
				$('#DadosUser__Username').html('').removeAttr('title');

			} else {

				var nomeShort = nome.length >= 12 ? nome.substr(0, 12) + '...' : nome;

				// Preencher dados do usuario na pagina
				$('#DadosUser__Username').html(nomeShort).attr('title', nome);

			}

			hideLoaderGeneric();

		}, 100);
	});
});