/*
* Example DriverPg C++
* @package     main
* @author      @jeffotoni
* @size        10/09/2018
*/

#include <iostream>

// ex: c++
class DriverPg
{
    private:
        
        /* este objeto sera nossa 
        instancia estatica */
        static DriverPg* instance;

        /* Construtor de nossa instancia */
        DriverPg();

    public:
    	int port;
        /* Metodo static public */
        static DriverPg* Connect();
};

/* Null, porque a instancia ira 
inciar conforme for chamada. */
DriverPg* DriverPg::instance = 0;

// nosso método 
DriverPg* DriverPg::Connect()
{
    if (instance == 0)
    {
        instance = new DriverPg();
        instance->port = 5432;
    }

    return instance;
}

DriverPg::DriverPg()
{}

int main()
{
    //new DriverPg(); // nao irá funcionar
    DriverPg* s = DriverPg::Connect(); // Ok
    DriverPg* r = DriverPg::Connect();
    int p = DriverPg::Connect()->port;

    // objetos retornados
    std::cout << s << std::endl;
    std::cout << r << std::endl;
    std::cout << p << std::endl;
}