package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraFlashControl struct {
	QMediaControl
}

type QCameraFlashControl_ITF interface {
	QMediaControl_ITF
	QCameraFlashControl_PTR() *QCameraFlashControl
}

func PointerFromQCameraFlashControl(ptr QCameraFlashControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraFlashControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraFlashControlFromPointer(ptr unsafe.Pointer) *QCameraFlashControl {
	var n = new(QCameraFlashControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraFlashControl_") {
		n.SetObjectName("QCameraFlashControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraFlashControl) QCameraFlashControl_PTR() *QCameraFlashControl {
	return ptr
}

func (ptr *QCameraFlashControl) FlashMode() QCameraExposure__FlashMode {
	defer qt.Recovering("QCameraFlashControl::flashMode")

	if ptr.Pointer() != nil {
		return QCameraExposure__FlashMode(C.QCameraFlashControl_FlashMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFlashControl) ConnectFlashReady(f func(ready bool)) {
	defer qt.Recovering("connect QCameraFlashControl::flashReady")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_ConnectFlashReady(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flashReady", f)
	}
}

func (ptr *QCameraFlashControl) DisconnectFlashReady() {
	defer qt.Recovering("disconnect QCameraFlashControl::flashReady")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_DisconnectFlashReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flashReady")
	}
}

//export callbackQCameraFlashControlFlashReady
func callbackQCameraFlashControlFlashReady(ptr unsafe.Pointer, ptrName *C.char, ready C.int) {
	defer qt.Recovering("callback QCameraFlashControl::flashReady")

	if signal := qt.GetSignal(C.GoString(ptrName), "flashReady"); signal != nil {
		signal.(func(bool))(int(ready) != 0)
	}

}

func (ptr *QCameraFlashControl) FlashReady(ready bool) {
	defer qt.Recovering("QCameraFlashControl::flashReady")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_FlashReady(ptr.Pointer(), C.int(qt.GoBoolToInt(ready)))
	}
}

func (ptr *QCameraFlashControl) IsFlashModeSupported(mode QCameraExposure__FlashMode) bool {
	defer qt.Recovering("QCameraFlashControl::isFlashModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraFlashControl_IsFlashModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFlashControl) IsFlashReady() bool {
	defer qt.Recovering("QCameraFlashControl::isFlashReady")

	if ptr.Pointer() != nil {
		return C.QCameraFlashControl_IsFlashReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraFlashControl) SetFlashMode(mode QCameraExposure__FlashMode) {
	defer qt.Recovering("QCameraFlashControl::setFlashMode")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_SetFlashMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFlashControl) DestroyQCameraFlashControl() {
	defer qt.Recovering("QCameraFlashControl::~QCameraFlashControl")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_DestroyQCameraFlashControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraFlashControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraFlashControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraFlashControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraFlashControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraFlashControlTimerEvent
func callbackQCameraFlashControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraFlashControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraFlashControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraFlashControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraFlashControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraFlashControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraFlashControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraFlashControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraFlashControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraFlashControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraFlashControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraFlashControlChildEvent
func callbackQCameraFlashControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraFlashControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraFlashControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraFlashControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraFlashControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraFlashControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraFlashControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraFlashControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraFlashControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraFlashControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraFlashControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraFlashControlCustomEvent
func callbackQCameraFlashControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraFlashControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraFlashControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraFlashControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraFlashControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraFlashControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraFlashControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraFlashControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}