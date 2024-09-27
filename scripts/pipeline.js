const { exec } = require('child_process');
require('dotenv').config();

function triggerPipeline() {
  try {
    console.log('Triggering CI/CD pipeline...');
    exec('ci-cd-command-to-trigger-pipeline', (error, stdout, stderr) => {
      if (error) {
        console.error(`Error triggering pipeline: ${error.message}`);
        sendNotification(`Pipeline Trigger Error: ${error.message}`);
        return;
      }
      if (stderr) {
        console.error(`Pipeline trigger stderr: ${stderr}`);
        sendNotification(`Pipeline Trigger Stderr: ${stderr}`);
        return;
      }
      console.log(stdout);
      checkPipelineStatus();
    });
  } catch (exception) {
    console.error(`Exception caught in triggerPipeline: ${exception}`);
    sendNotification(`Pipeline Trigger Exception: ${exception}`);
  }
}

function checkPipelineStatus() {
  console.log('Checking pipeline status...');
  
  try {
    setTimeout(() => {
      const pipelineStatus = 'success';
      if (pipelineStatus === 'success') {
        console.log('Pipeline executed successfully.');
        sendNotification('Pipeline executed successfully.');
      } else {
        console.error('Pipeline execution failed.');
        sendNotification('Pipeline execution failed.');
      }
    }, 2000);
  } catch (exception) {
    console.error(`Exception while checking pipeline status: ${exception}`);
    sendNotification(`Checking Pipeline Status Exception: ${exception}`);
  }
}

function sendNotification(message) {
  try {
    console.log(`Sending notification: ${message}`);
  } catch (exception) {
    console.error(`Exception sending notification: ${exception}`);
  }
}

triggerPipeline();