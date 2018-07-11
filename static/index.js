console.log('abcd');
console.log($('body').html());

const socket = io('http://localhost:5000/socket.io');
// console.log(socket);
// jQuery('#sendMsg').html('abcd');


// $('body').append('<p>dupa</p>')
socket.on('connect', function(){
  console.log('connected');
  socket.emit('newUserJoined', 'Testowy user');


$('#formArea').on('submit',function(e){
  e.preventDefault();
  console.log('Submitted');
})

$("#sendMsg").click(function(){
  console.log('Clicked');
  let messageString = $('#messageValue').val();
  // if(messageString.trim().length === 0){alert('Cannot send empty message!');return}
  socket.emit('newMessage', messageString.trim());
  appendToDiv(messageString.trim());
  $('#messageValue').val('');

});


// socket.emit('hello','adadijasijaisjdsiaj');
socket.on('newServerMessage', function(element){
  console.log('Dupa');
  console.log(element);
  appendToDiv(element);
})

socket.on('utilMessage', function(utilMessage){
  console.log('Util');
  console.log(utilMessage);
  appendToDiv(utilMessage);
})

})


function appendToDiv(messageText){
  $('#chatWindow').append(`<p>${messageText}</p>`);
}
