// Package CommonExit provides a tiny project convention for writing
// single-pass blocks that have one common exit point.
//
// The package intentionally contains one exported value: Block.
// It is designed for codebases that prefer a visible common-exit style
// over multiple early returns, especially while refactoring staged logic.
package CommonExit

// Block marks a single-pass block with a shared exit point.
//
// Use it with `for range` when a function benefits from:
//   - staged logic,
//   - early bailout using break,
//   - one final return statement,
//   - one common cleanup/logging/error-handling point,
//   - a visual marker that the block may later be extracted into a function.
//
// Example:
//
// 	var err error
// 	for range CommonExit.Block {
// 		err = validate()
// 		if err != nil {
// 			break
// 		}
//
// 		err = execute()
// 		if err != nil {
// 			break
// 		}
// 	}
// 	return err
//
// Do not use Block for actual looping. It exists to make a deliberately
// single-pass flow-control block explicit to human readers.
var Block = [...]struct{}{{}}
