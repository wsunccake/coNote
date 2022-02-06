package chapter15.example;

import javafx.application.Application;
import javafx.event.ActionEvent;
import javafx.event.EventHandler;
import javafx.geometry.Pos;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.layout.HBox;
import javafx.stage.Stage;

public class AnonymousHandlerDemo extends Application{
    @Override
    public void start(Stage primaryStage) throws Exception {
        HBox hBox = new HBox();
        hBox.setSpacing(10);
        hBox.setAlignment(Pos.CENTER);
        Button btNew = new Button("New");
        Button btOpen = new Button("Open");
        Button btSave = new Button("Save");
        Button btPrint = new Button("Print");
        Button btQuit = new Button("Quit");
        hBox.getChildren().addAll(btNew, btOpen, btPrint, btSave, btQuit);

        NewHandleClass handle1 = new NewHandleClass();
        btNew.setOnAction(handle1);

        OpenHandleClass handle2 = new OpenHandleClass();
        btOpen.setOnAction(handle2);

        btSave.setOnAction(new SaveHandleClass());

        btPrint.setOnAction(new EventHandler< ActionEvent>() {
            @Override
            public void handle(ActionEvent e) {
                System.out.println("Process Print");
            }
        });

        btQuit.setOnAction((ActionEvent event) -> {
            System.out.println("Process Quit");
        });

        Scene scene = new Scene(hBox, 300, 50);
        primaryStage.setTitle("AnonymousHandlerDemo");
        primaryStage.setScene(scene);
        primaryStage.show();
    }

    class OpenHandleClass implements EventHandler<ActionEvent> {
        @Override
        public void handle(ActionEvent event) {
            System.out.println("Process Open");
        }
    }

    class SaveHandleClass implements EventHandler<ActionEvent> {
        @Override
        public void handle(ActionEvent event) {
            System.out.println("Process Save");
        }
    }
}

class NewHandleClass implements EventHandler<ActionEvent> {
    @Override
    public void handle(ActionEvent event) {
        System.out.println("Process New");
    }
}