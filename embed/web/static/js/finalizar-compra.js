$(document).ready(function() {

    showLoaderGeneric(function() {

        setTimeout(function() {

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

            var checkUser = (nome !== null && nome.length > 0) &&
                (cep !== null && cep.length > 0) &&
                (uf !== null && uf.length > 0) &&
                (cidade !== null && cidade.length > 0) &&
                (bairro !== null && bairro.length > 0) &&
                (logradouro !== null && logradouro.length > 0) &&
                (numero !== null && numero.length > 0 && numero > 0);


            // Caso ja tenha dados do usuario, 
            // redirecionar para o finalizar compra

            if (!checkUser) {

                window.location.href = 'index.html';

            }

            var nomeShort = nome.length >= 12 ? nome.substr(0, 12) + '...' : nome;

            complemento = complemento ? ' / ' + complemento : '';

            // Preencher dados do usuario na pagina
            $('#DadosUser__Username').html(nomeShort).attr('title', nome);
            $('#DadosUser__Nome').html(nome);
            $('#DadosUser__Cep').html(cep);
            $('#DadosUser__Uf').html(uf);
            $('#DadosUser__Cidade').html(cidade);
            $('#DadosUser__Bairro').html(bairro);
            $('#DadosUser__Logradouro').html(logradouro);
            $('#DadosUser__Numero').html(numero);
            $('#DadosUser__Complemento').html(complemento);

            hideLoaderGeneric();

        }, 100);
    });


    // Mascara para os campos de Numero do endereco
    $('.form-control.form-control-cartao').formatter({
        'pattern': '{{9999}}',
        'persistent': false
    });

    // Mascara para os campos de Data validade
    $('.form-control.form-control-validade').formatter({
        'pattern': '{{99}}/{{9999}}',
        'persistent': false
    });

    // Mascara para os campos de CPF
    $('.form-control.form-control-cpf').formatter({
        'pattern': '{{999}}.{{999}}.{{999}}-{{99}}',
        'persistent': false
    });

    // Mascara para os campos de CVV
    $('.form-control.form-control-cvv').formatter({
        'pattern': '{{999}}',
        'persistent': false
    });

});



// Selecionar forma de pagamento: boleto x cartão
$(document).on('change', '.rdb-forma-pagamento', function() {
    var forma = $(this).attr('data-forma');

    $('.opcao-forma-pagamento').removeClass('opcao-selected');
    $('.opcao-fp-' + forma).addClass('opcao-selected');

    if (forma == 'boleto') {

        $('#FormaPagamentoNome').html('Boleto Bancário');
        $('#FormaPagamentoDescricao').html('');
        $('#FormaPagamentoValor').html('(10% OFF)');

        $('#PedidoValorTotal').html('R$ 792,00');

    } else {

        $('#FormaPagamentoNome').html('Cart&atilde;o de Cr&eacute;dito');
        $('#FormaPagamentoDescricao').html('');
        $('#FormaPagamentoValor').html('&Agrave; vista');

        $('#txtCartaoParcelas').val(1);

        $('#PedidoValorTotal').html('R$ 880,00');

    }
});



const valorParcelas = {
    1: { valor: 880.0, descricao: 'À vista' },
    2: { valor: 440.0, descricao: '2x de R$ 440,00 (sem juros)' },
    3: { valor: 293.34, descricao: '3x de R$ 293,34 (sem juros)' },
    4: { valor: 220.0, descricao: '4x de R$ 220,00 (sem juros)' },
    5: { valor: 176.0, descricao: '5x de R$ 176,00 (sem juros)' },
    6: { valor: 146.67, descricao: '6x de R$ 146,67 (sem juros)' },
    7: { valor: 125.72, descricao: '7x de R$ 125,72 (sem juros)' },
    8: { valor: 110.0, descricao: '8x de R$ 110,00 (sem juros)' }
};



// Selecionar parcela do cartao
$(document).on('change', '#txtCartaoParcelas', function() {
    var valorParcela = $(this).val();

    $('#FormaPagamentoValor').html(valorParcelas[valorParcela].descricao);
});



// Preenchimento do numero do cartao
$(document).on('keyup', '#txtCartaoNumero4', function() {
    var numero4 = $(this).val();

    if (numero4.length >= 3) {

        $('#FormaPagamentoDescricao').html('final ' + numero4);

    } else {

        $('#FormaPagamentoDescricao').html('');

    }
});



// Selecionar frete
$(document).on('change', '.rdb-frete', function() {
    var valor = $(this).attr('data-valor');
    var dias = $(this).attr('data-dias');

    $('#FreteNome').html($(this).val());
    $('#FreteValor').html(valor);
    $('#FreteDias').html(dias);
});


// Finalizar compra e efetuar pagamento, enviar JSON no POST
$(document).on('click', '.btn-efetuar-pagamento', function() {

    var isValid = true;

    var param = {};

    // Dados de nome e endereco 
    param.nome = LocalStorage.get('gcloud.poc.user_nome');
    param.endereco_cep = LocalStorage.get('gcloud.poc.endereco_cep');
    param.endereco_uf = LocalStorage.get('gcloud.poc.endereco_uf');
    param.endereco_cidade = LocalStorage.get('gcloud.poc.endereco_cidade');
    param.endereco_bairro = LocalStorage.get('gcloud.poc.endereco_bairro');
    param.endereco_logradouro = LocalStorage.get('gcloud.poc.endereco_logradouro');
    param.endereco_numero = LocalStorage.get('gcloud.poc.endereco_numero');
    param.endereco_complemento = LocalStorage.get('gcloud.poc.endereco_complemento');

    param.endereco_numero = Number.parseInt(param.endereco_numero);


    // Itens / Produtos do carrinho
    param.items = [{
            codigo: 'COD4623784',
            nome: "Camisa Docker e Go",
            qtd: 1,
            preco: 79.0
        },
        {
            codigo: 'CODD2344567',
            nome: "Camisa Zombie Go",
            qtd: 2,
            preco: 168.0
        },
        {
            codigo: 'CODD9871237',
            nome: "Camisa Go",
            qtd: 4,
            preco: 356.0
        },
        {
            codigo: 'CODD6547893',
            nome: "Camisa Dragon Ball Go",
            qtd: 3,
            preco: 277.0
        }
    ];


    // Dados do frente
    $('.rdb-frete').each(function() {
        if ($(this).is(':checked')) {
            param.entrega_nome = $(this).val();
            param.entrega_dias = Number.parseInt($(this).attr('data-dias'));
            param.entrega_valor = Number.parseFloat($(this).attr('data-valor'));
        }
    });



    // Dados da forma de pagamento: boleto / cartao

    var forma_pagamento = '';
    $('.rdb-forma-pagamento').each(function() {
        if ($(this).is(':checked')) {
            forma_pagamento = $(this).val();
        }
    });

    if (forma_pagamento == 'boleto') {

        param.forma_pagamento = 'boleto';
        param.boleto_desconto_porcentagem = 10;
        param.boleto_desconto_valor = 0;
        param.boleto_valor = 792.0;

    } else {

        param.forma_pagamento = 'cartao';

        var numero1 = $('#txtCartaoNumero1').val().trim();
        var numero2 = $('#txtCartaoNumero2').val().trim();
        var numero3 = $('#txtCartaoNumero3').val().trim();
        var numero4 = $('#txtCartaoNumero4').val().trim();
        var titular = $('#txtCartaoTitular').val().trim();
        var cpf = $('#txtCartaoCpf').val().trim();
        var validade = $('#txtCartaoValidade').val().trim();
        var cvv = $('#txtCartaoCvv').val().trim();
        var parcelas = $('#txtCartaoParcelas').val().trim();

        parcelas = Number.parseInt(parcelas);


        if (!numero1 || numero1.length == 0) {

            $('#txtCartaoNumero1').addClass('form-control-danger');
            $('#txtCartaoNumero1').parent('.form-group').addClass('has-danger');

            isValid = false;

        } else {
            $('#txtCartaoNumero1').removeClass('form-control-danger');
            $('#txtCartaoNumero1').parent('.form-group').removeClass('has-danger');
        }

        if (!numero2 || numero2.length == 0) {

            $('#txtCartaoNumero2').addClass('form-control-danger');
            $('#txtCartaoNumero2').parent('.form-group').addClass('has-danger');

            isValid = false;

        } else {
            $('#txtCartaoNumero2').removeClass('form-control-danger');
            $('#txtCartaoNumero2').parent('.form-group').removeClass('has-danger');
        }

        if (!numero3 || numero3.length == 0) {

            $('#txtCartaoNumero3').addClass('form-control-danger');
            $('#txtCartaoNumero3').parent('.form-group').addClass('has-danger');

            isValid = false;

        } else {
            $('#txtCartaoNumero3').removeClass('form-control-danger');
            $('#txtCartaoNumero3').parent('.form-group').removeClass('has-danger');
        }

        if (!numero4 || numero4.length == 0) {

            $('#txtCartaoNumero4').addClass('form-control-danger');
            $('#txtCartaoNumero4').parent('.form-group').addClass('has-danger');

            isValid = false;

        } else {
            $('#txtCartaoNumero4').removeClass('form-control-danger');
            $('#txtCartaoNumero4').parent('.form-group').removeClass('has-danger');
        }

        if (!titular || titular.length == 0) {

            $('#txtCartaoTitular').addClass('form-control-danger');
            $('#txtCartaoTitular').parent('.form-group').addClass('has-danger');

            isValid = false;

        } else {
            $('#txtCartaoTitular').removeClass('form-control-danger');
            $('#txtCartaoTitular').parent('.form-group').removeClass('has-danger');
        }

        if (!cpf || cpf.length == 0) {

            $('#txtCartaoCpf').addClass('form-control-danger');
            $('#txtCartaoCpf').parent('.form-group').addClass('has-danger');

            isValid = false;

        } else {
            $('#txtCartaoCpf').removeClass('form-control-danger');
            $('#txtCartaoCpf').parent('.form-group').removeClass('has-danger');
        }

        if (!validade || validade.length == 0) {

            $('#txtCartaoValidade').addClass('form-control-danger');
            $('#txtCartaoValidade').parent('.form-group').addClass('has-danger');

            isValid = false;

        } else {
            $('#txtCartaoValidade').removeClass('form-control-danger');
            $('#txtCartaoValidade').parent('.form-group').removeClass('has-danger');
        }

        if (!cvv || cvv.length == 0) {

            $('#txtCartaoCvv').addClass('form-control-danger');
            $('#txtCartaoCvv').parent('div').parent('div').parent('.form-group').addClass('has-danger');

            isValid = false;

        } else {
            $('#txtCartaoCvv').removeClass('form-control-danger');
            $('#txtCartaoCvv').parent('div').parent('div').parent('.form-group').removeClass('has-danger');
        }

        if (!parcelas || parcelas.length == 0) {

            $('#txtCartaoParcelas').addClass('form-control-danger');
            $('#txtCartaoParcelas').parent('.form-group').addClass('has-danger');

            isValid = false;

        } else {
            $('#txtCartaoParcelas').removeClass('form-control-danger');
            $('#txtCartaoParcelas').parent('.form-group').removeClass('has-danger');
        }


        if (!isValid) {

            toastr.error('Por favor, preencha os campos indicados!');

        } else {

            var numeroCartao = numero1 + numero2 + numero3 + numero4;

            var ano, mes;
            var validadeParts = validade.split('/');

            if (validadeParts[0] !== null && validadeParts[0] !== undefined) {
                mes = validadeParts[0];
            }

            if (validadeParts[1] !== null && validadeParts[1] !== undefined) {
                ano = validadeParts[1];
            }

            param.forma_pagamento = forma_pagamento;

            param.cartao_bandeira = 'Mastercard';

            param.cartao_numero = numeroCartao;
            param.cartao_nome_titular = titular;
            param.cartao_cpf = cpf;
            param.cartao_mes = Number.parseInt(mes);
            param.cartao_ano = Number.parseInt(ano);
            param.cartao_cvv = cvv;

            param.cartao_parcela_qtd = parcelas;
            param.cartao_parcela_valor = valorParcelas[parcelas].valor;
            param.cartao_parcela_descricao = valorParcelas[parcelas].descricao;

        }

    }

    console.log(param);

    if (isValid) {

        var path = 'http://localhost:8080/order';

        var objHeaders = {
            //'KEY': '$%28*10393948.@2019'
            // "Access-Control-Allow-Origin": "*",
            // "Access-Control-Allow-Headers": "Authorization", 
            // "Access-Control-Allow-Methods": "GET, POST, OPTIONS, PUT, PATCH, DELETE" ,
            // "Content-Type": "application/json;charset=UTF-8" 
        };

        showLoaderButton($('.btn-efetuar-pagamento'), function() {
            AxiosJson('POST', path, param, objHeaders, function(data, status) {

                console.log('Status > ', status);
                console.log('Data > ', data);

                if (status != 200) {

                    if (data.msg !== null && data.msg !== undefined) {


                        toastr.error(data.msg);

                        // $('.alerta-finalizar-erro .alerta-mensagem').html(data.msg);

                        $('.alerta-finalizar').hide();
                        // $('.alerta-finalizar-erro').slideDown();


                    } else {
                        toastr.error("Algo de estranho ocorreu em seu res.msg!");
                    }

                } else {

                    if (data.idpedido !== null && data.idpedido !== undefined) {

                        $('#AlertaUserNome').html(param.nome);
                        $('#AlertaNumeroPedido').html('#' + data.idpedido);

                        $('.alerta-finalizar').hide();
                        $('.alerta-finalizar-sucesso').slideDown();

                    }

                }

                hideLoaderButton($('.btn-efetuar-pagamento'));

            });
        });

    }
});


$(document).on('click', '.alerta-finalizar .btn-ok-close', function() {
    $('.alerta-finalizar').hide();

    $('#AlertaUserNome').html('');
    $('#AlertaNumeroPedido').html('');

    $('.alerta-finalizar-erro .alerta-mensagem').html('');
});