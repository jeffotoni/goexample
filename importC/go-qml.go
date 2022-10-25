func main() {
        err := qml.Run(run)
        ...
}

func run() error {
        engine := qml.NewEngine()
        component, err := engine.LoadFile("file.qml")
        if err != nil {
                return err
        }
        win := component.CreateWindow(nil)
        win.Show()
        win.Wait()
        return nil
}