package bitcoin

// WARNING!!!! This should not be a float!! float is probably not safe enough!
// TODO: create safe Amount struct with json marshalling and pretty printing.
type Amount float64
