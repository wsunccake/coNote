package chapter15.example;

import javafx.application.Application;
import javafx.geometry.Pos;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.input.KeyCode;
import javafx.scene.input.MouseButton;
import javafx.scene.layout.BorderPane;
import javafx.scene.layout.HBox;
import javafx.stage.Stage;

public class ControlCircleWithMouseAndKey extends Application {
    private CirclePane circlePane = new CirclePane();

    @Override
    public void start(Stage primaryStage) throws Exception {
        HBox hBox = new HBox();
        hBox.setSpacing(10);
        hBox.setAlignment(Pos.CENTER);
        Button btEnlarge = new Button("Enlarge");
        Button btShrink = new Button("Shrink");
        hBox.getChildren().add(btEnlarge);
        hBox.getChildren().add(btShrink);

        btEnlarge.setOnAction(event -> circlePane.enlarge());
        btShrink.setOnAction(event -> circlePane.shrink());

        circlePane.setOnMouseClicked(event -> {
            if (event.getButton() == MouseButton.PRIMARY) {
                circlePane.enlarge();
            } else if (event.getButton() == MouseButton.SECONDARY) {
                circlePane.shrink();
            }
        });

        circlePane.setOnKeyPressed(event -> {
            if (event.getCode() == KeyCode.U) {
                circlePane.enlarge();
            } else if (event.getCode() == KeyCode.D) {
                circlePane.shrink();
            }
        });

        BorderPane borderPane = new BorderPane();
        borderPane.setCenter(circlePane);
        borderPane.setBottom(hBox);
        BorderPane.setAlignment(hBox, Pos.CENTER);

        Scene scene = new Scene(borderPane, 200, 150);
        primaryStage.setTitle("ControlCircle");
        primaryStage.setScene(scene);
        primaryStage.show();

        circlePane.requestFocus();
    }
}
