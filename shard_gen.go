// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package dagstore

import (
	"fmt"
	"io"
	"math"

	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

func (t *PersistedShard) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{165}); err != nil {
		return err
	}

	// t.Key (string) (string)
	if len("Key") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Key\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Key")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Key")); err != nil {
		return err
	}

	if len(t.Key) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Key was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.Key)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.Key)); err != nil {
		return err
	}

	// t.URL (string) (string)
	if len("URL") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"URL\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("URL")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("URL")); err != nil {
		return err
	}

	if len(t.URL) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.URL was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.URL)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.URL)); err != nil {
		return err
	}

	// t.State (dagstore.ShardState) (uint8)
	if len("State") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"State\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("State")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("State")); err != nil {
		return err
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.State))); err != nil {
		return err
	}

	// t.Indexed (bool) (bool)
	if len("Indexed") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Indexed\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Indexed")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Indexed")); err != nil {
		return err
	}

	// t.TransientPath (string) (string)
	if len("TransientPath") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"TransientPath\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("TransientPath")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("TransientPath")); err != nil {
		return err
	}

	if len(t.TransientPath) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.TransientPath was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.TransientPath)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.TransientPath)); err != nil {
		return err
	}
	return nil
}

func (t *PersistedShard) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("PersistedShard: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(br)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Key (string) (string)
		case "Key":

			{
				sval, err := cbg.ReadString(br)
				if err != nil {
					return err
				}

				t.Key = string(sval)
			}
			// t.URL (string) (string)
		case "URL":

			{
				sval, err := cbg.ReadString(br)
				if err != nil {
					return err
				}

				t.URL = string(sval)
			}
			// t.State (dagstore.ShardState) (uint8)
		case "State":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}
			if maj != cbg.MajUnsignedInt {
				return fmt.Errorf("wrong type for uint8 field")
			}
			if extra > math.MaxUint8 {
				return fmt.Errorf("integer in input was too large for uint8 field")
			}
			t.State = ShardState(extra)
			// t.Indexed (bool) (bool)
		case "Indexed":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}

			// t.TransientPath (string) (string)
		case "TransientPath":

			{
				sval, err := cbg.ReadString(br)
				if err != nil {
					return err
				}

				t.TransientPath = string(sval)
			}

		default:
			return fmt.Errorf("unknown struct field %d: '%s'", i, name)
		}
	}

	return nil
}
