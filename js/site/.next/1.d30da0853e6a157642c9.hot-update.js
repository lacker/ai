webpackHotUpdate(1,{

/***/ 139:
/***/ function(module, exports, __webpack_require__) {

	'use strict';
	
	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	
	var _stringify = __webpack_require__(89);
	
	var _stringify2 = _interopRequireDefault(_stringify);
	
	var _getPrototypeOf = __webpack_require__(21);
	
	var _getPrototypeOf2 = _interopRequireDefault(_getPrototypeOf);
	
	var _classCallCheck2 = __webpack_require__(22);
	
	var _classCallCheck3 = _interopRequireDefault(_classCallCheck2);
	
	var _createClass2 = __webpack_require__(23);
	
	var _createClass3 = _interopRequireDefault(_createClass2);
	
	var _possibleConstructorReturn2 = __webpack_require__(25);
	
	var _possibleConstructorReturn3 = _interopRequireDefault(_possibleConstructorReturn2);
	
	var _inherits2 = __webpack_require__(24);
	
	var _inherits3 = _interopRequireDefault(_inherits2);
	
	var _react = __webpack_require__(20);
	
	var _react2 = _interopRequireDefault(_react);
	
	var _mobxReact = __webpack_require__(137);
	
	var _ChatInput = __webpack_require__(138);
	
	var _ChatInput2 = _interopRequireDefault(_ChatInput);
	
	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }
	
	function makeID() {
	  var answer = '';
	  for (var i = 0; i < 8; i++) {
	    var index = Math.floor(Math.random() * 32);
	    answer += 'ABCDEFGHJKMNPQRSTVWXYZ0123456789'[index];
	  }
	  return answer;
	}
	
	var ListView = function (_React$Component) {
	  (0, _inherits3.default)(ListView, _React$Component);
	
	  function ListView(props) {
	    (0, _classCallCheck3.default)(this, ListView);
	
	    var _this = (0, _possibleConstructorReturn3.default)(this, (ListView.__proto__ || (0, _getPrototypeOf2.default)(ListView)).call(this, props));
	
	    _this.handleSubmit = _this.handleSubmit.bind(_this);
	    return _this;
	  }
	
	  (0, _createClass3.default)(ListView, [{
	    key: 'handleSubmit',
	    value: function handleSubmit(value) {
	      console.log('at the start of handleSubmit messages is:', (0, _stringify2.default)(this.props.messages));
	
	      // TODO: use chat server instead of just locally doing stuff
	      // TODO: figure out how to make this rerender the ListView
	      this.props.messages.push({
	        id: makeID(),
	        content: value
	      });
	
	      console.log('at the end of handleSubmit messages is:', (0, _stringify2.default)(this.props.messages));
	    }
	  }, {
	    key: 'render',
	    value: function render() {
	      console.log('rendering messages:', (0, _stringify2.default)(this.props.messages));
	
	      return _react2.default.createElement(
	        'div',
	        null,
	        'Welcome to chat.',
	        _react2.default.createElement(
	          'ul',
	          null,
	          this.props.messages.map(function (message) {
	            return _react2.default.createElement(
	              'li',
	              { key: message.id },
	              message.content
	            );
	          })
	        ),
	        _react2.default.createElement(_ChatInput2.default, { onSubmit: this.handleSubmit })
	      );
	    }
	  }]);
	  return ListView;
	}(_react2.default.Component);
	
	ListView = (0, _mobxReact.observer)(ListView);
	
	exports.default = ListView;
	//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIkxpc3RWaWV3LmpzIl0sIm5hbWVzIjpbIm1ha2VJRCIsImFuc3dlciIsImkiLCJpbmRleCIsIk1hdGgiLCJmbG9vciIsInJhbmRvbSIsIkxpc3RWaWV3IiwicHJvcHMiLCJoYW5kbGVTdWJtaXQiLCJiaW5kIiwidmFsdWUiLCJjb25zb2xlIiwibG9nIiwibWVzc2FnZXMiLCJwdXNoIiwiaWQiLCJjb250ZW50IiwibWFwIiwibWVzc2FnZSIsIkNvbXBvbmVudCJdLCJtYXBwaW5ncyI6Ijs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUE7Ozs7QUFDQTs7QUFFQTs7Ozs7O0FBRUEsU0FBU0EsTUFBVCxHQUFrQjtBQUNoQixNQUFJQyxTQUFTLEVBQWI7QUFDQSxPQUFLLElBQUlDLElBQUksQ0FBYixFQUFnQkEsSUFBSSxDQUFwQixFQUF1QkEsR0FBdkIsRUFBNEI7QUFDMUIsUUFBSUMsUUFBUUMsS0FBS0MsS0FBTCxDQUFXRCxLQUFLRSxNQUFMLEtBQWdCLEVBQTNCLENBQVo7QUFDQUwsY0FBVSxtQ0FBbUNFLEtBQW5DLENBQVY7QUFDRDtBQUNELFNBQU9GLE1BQVA7QUFDRDs7SUFFS00sUTs7O0FBQ0osb0JBQVlDLEtBQVosRUFBbUI7QUFBQTs7QUFBQSwwSUFDWEEsS0FEVzs7QUFHakIsVUFBS0MsWUFBTCxHQUFvQixNQUFLQSxZQUFMLENBQWtCQyxJQUFsQixPQUFwQjtBQUhpQjtBQUlsQjs7OztpQ0FFWUMsSyxFQUFPO0FBQ2xCQyxjQUFRQyxHQUFSLENBQ0UsMkNBREYsRUFFRSx5QkFBZSxLQUFLTCxLQUFMLENBQVdNLFFBQTFCLENBRkY7O0FBSUE7QUFDQTtBQUNBLFdBQUtOLEtBQUwsQ0FBV00sUUFBWCxDQUFvQkMsSUFBcEIsQ0FBeUI7QUFDdkJDLFlBQUloQixRQURtQjtBQUV2QmlCLGlCQUFTTjtBQUZjLE9BQXpCOztBQUtBQyxjQUFRQyxHQUFSLENBQ0UseUNBREYsRUFFRSx5QkFBZSxLQUFLTCxLQUFMLENBQVdNLFFBQTFCLENBRkY7QUFHRDs7OzZCQUVRO0FBQ1BGLGNBQVFDLEdBQVIsQ0FDRSxxQkFERixFQUVFLHlCQUFlLEtBQUtMLEtBQUwsQ0FBV00sUUFBMUIsQ0FGRjs7QUFJQSxhQUNFO0FBQUE7QUFBQTtBQUFBO0FBRUU7QUFBQTtBQUFBO0FBQ0csZUFBS04sS0FBTCxDQUFXTSxRQUFYLENBQW9CSSxHQUFwQixDQUF3QjtBQUFBLG1CQUN2QjtBQUFBO0FBQUEsZ0JBQUksS0FBS0MsUUFBUUgsRUFBakI7QUFBc0JHLHNCQUFRRjtBQUE5QixhQUR1QjtBQUFBLFdBQXhCO0FBREgsU0FGRjtBQU9FLDZEQUFXLFVBQVUsS0FBS1IsWUFBMUI7QUFQRixPQURGO0FBV0Q7OztFQXhDb0IsZ0JBQU1XLFM7O0FBMEM3QmIsV0FBVyx5QkFBU0EsUUFBVCxDQUFYOztrQkFFZUEsUSIsImZpbGUiOiJMaXN0Vmlldy5qcyIsInNvdXJjZVJvb3QiOiIvVXNlcnMvbGFja2VyL2FpL2pzL3NpdGUiLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgUmVhY3QgZnJvbSAncmVhY3QnO1xuaW1wb3J0IHsgb2JzZXJ2ZXIgfSBmcm9tICdtb2J4LXJlYWN0JztcblxuaW1wb3J0IENoYXRJbnB1dCBmcm9tICcuL0NoYXRJbnB1dCc7XG5cbmZ1bmN0aW9uIG1ha2VJRCgpIHtcbiAgbGV0IGFuc3dlciA9ICcnO1xuICBmb3IgKGxldCBpID0gMDsgaSA8IDg7IGkrKykge1xuICAgIGxldCBpbmRleCA9IE1hdGguZmxvb3IoTWF0aC5yYW5kb20oKSAqIDMyKTtcbiAgICBhbnN3ZXIgKz0gJ0FCQ0RFRkdISktNTlBRUlNUVldYWVowMTIzNDU2Nzg5J1tpbmRleF07XG4gIH1cbiAgcmV0dXJuIGFuc3dlcjtcbn1cblxuY2xhc3MgTGlzdFZpZXcgZXh0ZW5kcyBSZWFjdC5Db21wb25lbnQge1xuICBjb25zdHJ1Y3Rvcihwcm9wcykge1xuICAgIHN1cGVyKHByb3BzKTtcblxuICAgIHRoaXMuaGFuZGxlU3VibWl0ID0gdGhpcy5oYW5kbGVTdWJtaXQuYmluZCh0aGlzKTtcbiAgfVxuXG4gIGhhbmRsZVN1Ym1pdCh2YWx1ZSkge1xuICAgIGNvbnNvbGUubG9nKFxuICAgICAgJ2F0IHRoZSBzdGFydCBvZiBoYW5kbGVTdWJtaXQgbWVzc2FnZXMgaXM6JyxcbiAgICAgIEpTT04uc3RyaW5naWZ5KHRoaXMucHJvcHMubWVzc2FnZXMpKTtcblxuICAgIC8vIFRPRE86IHVzZSBjaGF0IHNlcnZlciBpbnN0ZWFkIG9mIGp1c3QgbG9jYWxseSBkb2luZyBzdHVmZlxuICAgIC8vIFRPRE86IGZpZ3VyZSBvdXQgaG93IHRvIG1ha2UgdGhpcyByZXJlbmRlciB0aGUgTGlzdFZpZXdcbiAgICB0aGlzLnByb3BzLm1lc3NhZ2VzLnB1c2goe1xuICAgICAgaWQ6IG1ha2VJRCgpLFxuICAgICAgY29udGVudDogdmFsdWUsXG4gICAgfSk7XG5cbiAgICBjb25zb2xlLmxvZyhcbiAgICAgICdhdCB0aGUgZW5kIG9mIGhhbmRsZVN1Ym1pdCBtZXNzYWdlcyBpczonLFxuICAgICAgSlNPTi5zdHJpbmdpZnkodGhpcy5wcm9wcy5tZXNzYWdlcykpO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIGNvbnNvbGUubG9nKFxuICAgICAgJ3JlbmRlcmluZyBtZXNzYWdlczonLFxuICAgICAgSlNPTi5zdHJpbmdpZnkodGhpcy5wcm9wcy5tZXNzYWdlcykpO1xuXG4gICAgcmV0dXJuIChcbiAgICAgIDxkaXY+XG4gICAgICBXZWxjb21lIHRvIGNoYXQuXG4gICAgICAgIDx1bD5cbiAgICAgICAgICB7dGhpcy5wcm9wcy5tZXNzYWdlcy5tYXAobWVzc2FnZSA9PiAoXG4gICAgICAgICAgICA8bGkga2V5PXttZXNzYWdlLmlkfT57bWVzc2FnZS5jb250ZW50fTwvbGk+XG4gICAgICAgICAgKSl9XG4gICAgICAgIDwvdWw+XG4gICAgICAgIDxDaGF0SW5wdXQgb25TdWJtaXQ9e3RoaXMuaGFuZGxlU3VibWl0fSAvPlxuICAgICAgPC9kaXY+XG4gICAgKTtcbiAgfVxufVxuTGlzdFZpZXcgPSBvYnNlcnZlcihMaXN0Vmlldyk7XG5cbmV4cG9ydCBkZWZhdWx0IExpc3RWaWV3O1xuIl19

/***/ }

})
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vLi9MaXN0Vmlldy5qcz8xMGExIl0sIm5hbWVzIjpbIm1ha2VJRCIsImFuc3dlciIsImkiLCJpbmRleCIsIk1hdGgiLCJmbG9vciIsInJhbmRvbSIsIkxpc3RWaWV3IiwicHJvcHMiLCJoYW5kbGVTdWJtaXQiLCJiaW5kIiwidmFsdWUiLCJjb25zb2xlIiwibG9nIiwibWVzc2FnZXMiLCJwdXNoIiwiaWQiLCJjb250ZW50IiwibWFwIiwibWVzc2FnZSIsIkNvbXBvbmVudCJdLCJtYXBwaW5ncyI6Ijs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQTs7OztBQUNBOztBQUVBOzs7Ozs7QUFFQSxVQUFTQSxNQUFULEdBQWtCO0FBQ2hCLE9BQUlDLFNBQVMsRUFBYjtBQUNBLFFBQUssSUFBSUMsSUFBSSxDQUFiLEVBQWdCQSxJQUFJLENBQXBCLEVBQXVCQSxHQUF2QixFQUE0QjtBQUMxQixTQUFJQyxRQUFRQyxLQUFLQyxLQUFMLENBQVdELEtBQUtFLE1BQUwsS0FBZ0IsRUFBM0IsQ0FBWjtBQUNBTCxlQUFVLG1DQUFtQ0UsS0FBbkMsQ0FBVjtBQUNEO0FBQ0QsVUFBT0YsTUFBUDtBQUNEOztLQUVLTSxROzs7QUFDSixxQkFBWUMsS0FBWixFQUFtQjtBQUFBOztBQUFBLDJJQUNYQSxLQURXOztBQUdqQixXQUFLQyxZQUFMLEdBQW9CLE1BQUtBLFlBQUwsQ0FBa0JDLElBQWxCLE9BQXBCO0FBSGlCO0FBSWxCOzs7O2tDQUVZQyxLLEVBQU87QUFDbEJDLGVBQVFDLEdBQVIsQ0FDRSwyQ0FERixFQUVFLHlCQUFlLEtBQUtMLEtBQUwsQ0FBV00sUUFBMUIsQ0FGRjs7QUFJQTtBQUNBO0FBQ0EsWUFBS04sS0FBTCxDQUFXTSxRQUFYLENBQW9CQyxJQUFwQixDQUF5QjtBQUN2QkMsYUFBSWhCLFFBRG1CO0FBRXZCaUIsa0JBQVNOO0FBRmMsUUFBekI7O0FBS0FDLGVBQVFDLEdBQVIsQ0FDRSx5Q0FERixFQUVFLHlCQUFlLEtBQUtMLEtBQUwsQ0FBV00sUUFBMUIsQ0FGRjtBQUdEOzs7OEJBRVE7QUFDUEYsZUFBUUMsR0FBUixDQUNFLHFCQURGLEVBRUUseUJBQWUsS0FBS0wsS0FBTCxDQUFXTSxRQUExQixDQUZGOztBQUlBLGNBQ0U7QUFBQTtBQUFBO0FBQUE7QUFFRTtBQUFBO0FBQUE7QUFDRyxnQkFBS04sS0FBTCxDQUFXTSxRQUFYLENBQW9CSSxHQUFwQixDQUF3QjtBQUFBLG9CQUN2QjtBQUFBO0FBQUEsaUJBQUksS0FBS0MsUUFBUUgsRUFBakI7QUFBc0JHLHVCQUFRRjtBQUE5QixjQUR1QjtBQUFBLFlBQXhCO0FBREgsVUFGRjtBQU9FLDhEQUFXLFVBQVUsS0FBS1IsWUFBMUI7QUFQRixRQURGO0FBV0Q7OztHQXhDb0IsZ0JBQU1XLFM7O0FBMEM3QmIsWUFBVyx5QkFBU0EsUUFBVCxDQUFYOzttQkFFZUEsUSIsImZpbGUiOiIxLmQzMGRhMDg1M2U2YTE1NzY0MmM5LmhvdC11cGRhdGUuanMiLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgUmVhY3QgZnJvbSAncmVhY3QnO1xuaW1wb3J0IHsgb2JzZXJ2ZXIgfSBmcm9tICdtb2J4LXJlYWN0JztcblxuaW1wb3J0IENoYXRJbnB1dCBmcm9tICcuL0NoYXRJbnB1dCc7XG5cbmZ1bmN0aW9uIG1ha2VJRCgpIHtcbiAgbGV0IGFuc3dlciA9ICcnO1xuICBmb3IgKGxldCBpID0gMDsgaSA8IDg7IGkrKykge1xuICAgIGxldCBpbmRleCA9IE1hdGguZmxvb3IoTWF0aC5yYW5kb20oKSAqIDMyKTtcbiAgICBhbnN3ZXIgKz0gJ0FCQ0RFRkdISktNTlBRUlNUVldYWVowMTIzNDU2Nzg5J1tpbmRleF07XG4gIH1cbiAgcmV0dXJuIGFuc3dlcjtcbn1cblxuY2xhc3MgTGlzdFZpZXcgZXh0ZW5kcyBSZWFjdC5Db21wb25lbnQge1xuICBjb25zdHJ1Y3Rvcihwcm9wcykge1xuICAgIHN1cGVyKHByb3BzKTtcblxuICAgIHRoaXMuaGFuZGxlU3VibWl0ID0gdGhpcy5oYW5kbGVTdWJtaXQuYmluZCh0aGlzKTtcbiAgfVxuXG4gIGhhbmRsZVN1Ym1pdCh2YWx1ZSkge1xuICAgIGNvbnNvbGUubG9nKFxuICAgICAgJ2F0IHRoZSBzdGFydCBvZiBoYW5kbGVTdWJtaXQgbWVzc2FnZXMgaXM6JyxcbiAgICAgIEpTT04uc3RyaW5naWZ5KHRoaXMucHJvcHMubWVzc2FnZXMpKTtcblxuICAgIC8vIFRPRE86IHVzZSBjaGF0IHNlcnZlciBpbnN0ZWFkIG9mIGp1c3QgbG9jYWxseSBkb2luZyBzdHVmZlxuICAgIC8vIFRPRE86IGZpZ3VyZSBvdXQgaG93IHRvIG1ha2UgdGhpcyByZXJlbmRlciB0aGUgTGlzdFZpZXdcbiAgICB0aGlzLnByb3BzLm1lc3NhZ2VzLnB1c2goe1xuICAgICAgaWQ6IG1ha2VJRCgpLFxuICAgICAgY29udGVudDogdmFsdWUsXG4gICAgfSk7XG5cbiAgICBjb25zb2xlLmxvZyhcbiAgICAgICdhdCB0aGUgZW5kIG9mIGhhbmRsZVN1Ym1pdCBtZXNzYWdlcyBpczonLFxuICAgICAgSlNPTi5zdHJpbmdpZnkodGhpcy5wcm9wcy5tZXNzYWdlcykpO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIGNvbnNvbGUubG9nKFxuICAgICAgJ3JlbmRlcmluZyBtZXNzYWdlczonLFxuICAgICAgSlNPTi5zdHJpbmdpZnkodGhpcy5wcm9wcy5tZXNzYWdlcykpO1xuXG4gICAgcmV0dXJuIChcbiAgICAgIDxkaXY+XG4gICAgICBXZWxjb21lIHRvIGNoYXQuXG4gICAgICAgIDx1bD5cbiAgICAgICAgICB7dGhpcy5wcm9wcy5tZXNzYWdlcy5tYXAobWVzc2FnZSA9PiAoXG4gICAgICAgICAgICA8bGkga2V5PXttZXNzYWdlLmlkfT57bWVzc2FnZS5jb250ZW50fTwvbGk+XG4gICAgICAgICAgKSl9XG4gICAgICAgIDwvdWw+XG4gICAgICAgIDxDaGF0SW5wdXQgb25TdWJtaXQ9e3RoaXMuaGFuZGxlU3VibWl0fSAvPlxuICAgICAgPC9kaXY+XG4gICAgKTtcbiAgfVxufVxuTGlzdFZpZXcgPSBvYnNlcnZlcihMaXN0Vmlldyk7XG5cbmV4cG9ydCBkZWZhdWx0IExpc3RWaWV3O1xuXG5cblxuLy8gV0VCUEFDSyBGT09URVIgLy9cbi8vIC4vTGlzdFZpZXcuanMiXSwic291cmNlUm9vdCI6IiJ9