using System;
using System.Collections.Generic;
using System.Linq;
using System.Runtime.InteropServices;
using System.Runtime.Remoting.Metadata.W3cXsd2001;
using System.Security.Cryptography.X509Certificates;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;


namespace ProyectoBingo
{
    /// <summary>
    /// Lógica de interacción para MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {
        public String user;
        public String localhost;
        public string readPort;
        public string writePort;
        public String gameMode;
        public String boardNumber;



        // Inicio
        public MainWindow()
        {
            InitializeComponent();
        }

        private void Button_Click(object sender, RoutedEventArgs e)
        {
            user = User.Text;
            localhost = Localhost.Text;
            readPort = ReadPort.Text;
            writePort = WritePort.Text;
            gameMode = Convert.ToString(GameMode.SelectedItem);
            boardNumber =Convert.ToString(BoardNumber.SelectedItem);

            /*MessageBox.Show(user);
            MessageBox.Show(Convert.ToString(localhost));
            MessageBox.Show(readPort);
            MessageBox.Show(writePort);
            MessageBox.Show(Convert.ToString(boardNumber));
            MessageBox.Show(Convert.ToString(gameMode));*/

              JuegoBingo Bingo = new JuegoBingo(localhost);
            this.Close();
            Bingo.Show();
        }
    }


}
