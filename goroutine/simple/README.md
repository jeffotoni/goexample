#Fan-out Goroutine


#### Goroutine e Pattern
Existem alguns patterns para concorrência, dentre eles: Fan-In e Fan-Out.

### Fan-Out:
Múltiplas funções podem ler do mesmo canal até que esse canal seja fechado, isso é chamado de fan-out . Isso fornece uma maneira de distribuir o trabalho entre um grupo de trabalhadores para paralelizar o uso da CPU e E/S.

### Fan-In:
É a forma de ler a partir de múltiplas entradas e prosseguir até que todas estejam fechadas, multiplexando os canais de entrada em um único canal que está fechado quando todas as entradas estão fechadas. Isso é chamado fan-in.