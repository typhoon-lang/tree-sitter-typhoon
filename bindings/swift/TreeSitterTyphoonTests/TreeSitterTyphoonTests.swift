import XCTest
import SwiftTreeSitter
import TreeSitterTyphoon

final class TreeSitterTyphoonTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_typhoon())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Typhoon grammar")
    }
}
