using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Shapes;
using System.Net.Http;
using System.Runtime.CompilerServices;
using System.Security.Cryptography.X509Certificates;


namespace ProyectoBingo
{

    // Clase Bola
    public class Ball
    {
        public string letter { get; set; }
        public int number { get; set; }

    }
    //Clase Casilla
    public class Tiles
    {
        public string letter { get; set; }
        public int number { get; set; }
        public bool taken { get; set; }
    }


    // Clase Tablero
    public class Board
    {
        public string name { get; set; }
        public Tiles[,,,,] Tile { get; set; }

    }

    //Clase GUI
    public class GUI
    {
        public Ball bolita { set; get; }
        public Board[,] cartones { set; get; }

        public string[] bingo { set; get; }
    }
    

    /// <summary>
    /// Lógica de interacción para JuegoBingo.xaml
    /// </summary>
    public partial class JuegoBingo : Window
    {
        HttpClient cliente = new HttpClient();


        public async Task Bingo(string localhost)
        {
            JuegoBingo Bingo = new JuegoBingo(); 
            await Bingo.Prueba();
            InitializeComponent();

        }

        private async Task Prueba()
        {
            HttpResponseMessage response = await cliente.GetAsync("https://jsonplaceholder.typicode.com/todos");
            Console.WriteLine(response);
        }
    }
    
}
