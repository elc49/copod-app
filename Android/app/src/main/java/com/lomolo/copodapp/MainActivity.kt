package com.lomolo.copodapp

import android.content.Intent
import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material3.Button
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.navigation.compose.rememberNavController
import com.lomolo.copodapp.di.appModule
import com.lomolo.copodapp.ui.theme.CopodAppTheme
import com.lomolo.copodapp.ui.viewmodels.GetDeviceDetails
import com.lomolo.copodapp.ui.viewmodels.MainViewModel
import org.koin.android.ext.koin.androidContext
import org.koin.androidx.viewmodel.ext.android.viewModel
import org.koin.core.context.startKoin

class MainActivity : ComponentActivity() {
    private val mainViewModel: MainViewModel by viewModel()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        startKoin {
            androidContext(this@MainActivity)
            modules(appModule)
        }
        enableEdgeToEdge()
        setContent {
            CopodAppTheme {
                Scaffold(modifier = Modifier.fillMaxSize()) { innerPadding ->
                    when (mainViewModel.gettingDeviceDetails) {
                        GetDeviceDetails.Success -> CopodApplication(
                            Modifier.padding(innerPadding),
                            rememberNavController(),
                            mainViewModel = mainViewModel,
                        )

                        GetDeviceDetails.Loading -> Column(
                            Modifier.fillMaxSize(),
                            verticalArrangement = Arrangement.Center,
                            horizontalAlignment = Alignment.CenterHorizontally,
                        ) {
                            CircularProgressIndicator(
                                Modifier.size(20.dp),
                            )
                        }

                        else -> Column(
                            Modifier.fillMaxSize(),
                            verticalArrangement = Arrangement.Center,
                            horizontalAlignment = Alignment.CenterHorizontally,
                        ) {
                            Text(
                                stringResource(R.string.something_wrong),
                                style = MaterialTheme.typography.titleMedium,
                                color = MaterialTheme.colorScheme.error,
                                textAlign = TextAlign.Center,
                            )
                            Button(
                                onClick = { mainViewModel.getDeviceDetails() },
                                shape = MaterialTheme.shapes.extraSmall,
                            ) {
                                Text(
                                    stringResource(R.string.retry),
                                    style = MaterialTheme.typography.titleMedium,
                                )
                            }
                        }
                    }
                }
            }
            mainViewModel.initialize()
        }
    }

    override fun onNewIntent(intent: Intent) {
        super.onNewIntent(intent)

        // Handle user signing in when app is active
        mainViewModel.setResultUrl(intent.data)
    }
}