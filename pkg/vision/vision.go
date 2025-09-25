package vision

// Camera defines the public interface for controlling a VISCA camera.
type Camera interface {
    ZoomIn(speed byte) error
    ZoomOut(speed byte) error
    ZoomStop() error
    FocusAuto() error
    FocusManual() error
    GetZoomPosition() (int, error)
    CheckAlive() error
    Close() error
}
