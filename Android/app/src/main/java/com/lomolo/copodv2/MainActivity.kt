package com.lomolo.copodv2

import android.content.Intent
import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Scaffold
import androidx.compose.ui.Modifier
import androidx.navigation.compose.rememberNavController
import com.lomolo.copodv2.di.appModule
import com.lomolo.copodv2.ui.theme.CopodV2Theme
import com.lomolo.copodv2.viewmodels.MainViewModel
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
            CopodV2Theme {
                Scaffold(modifier = Modifier.fillMaxSize()) { innerPadding ->
                    CopodApplication(
                        Modifier.padding(innerPadding),
                        rememberNavController(),
                        mainViewModel = mainViewModel,
                    )
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